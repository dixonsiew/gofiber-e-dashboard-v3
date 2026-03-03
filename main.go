package main

import (
    "edashboard/config"
    "edashboard/database"
    _ "edashboard/docs"
    "edashboard/router"
    "edashboard/utils"
    "errors"
    "fmt"
    "html/template"
    "os"
    "strings"
    "sync"

    swaggo "github.com/gofiber/contrib/v3/swaggo"
    fiberzerolog "github.com/gofiber/contrib/v3/zerolog"
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/compress"
    "github.com/gofiber/fiber/v3/middleware/cors"
    "github.com/gofiber/fiber/v3/middleware/healthcheck"
    "github.com/gofiber/fiber/v3/middleware/recover"
    "github.com/gofiber/fiber/v3/middleware/static"
)

// @title Swagger E-Dashboard API
// @version 1.0
// @description E-Dashboard API description.
// @BasePath /edashboard
func main() {
    defer utils.CatchPanic("main")
    runLogFile, _ := os.OpenFile(
        "app.log",
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0664,
    )
    defer runLogFile.Close()
    utils.SetLogger(runLogFile)
    ports := config.Config("port")
    lsport := strings.Split(ports, ",")
    app := fiber.New(fiber.Config{
        ErrorHandler: func(c fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            var e *fiber.Error
            if errors.As(err, &e) {
                code = e.Code
            }

            return c.Status(code).JSON(fiber.Map{
                "statusCode": code,
                "message":    err.Error(),
            })
        },
    })
    app.Use(recover.New())
    app.Use(compress.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
    }))
    app.Use(fiberzerolog.New(fiberzerolog.Config{
        Logger: &utils.Logger,
    }))
    database.InitDB()
    database.InitDBRs()
    defer database.CloseDB()
    defer database.CloseDBRs()

    basePath := "edashboard"
    initSwagger(app, basePath)
    app.Get(healthcheck.StartupEndpoint, healthcheck.New())
    app.Get(fmt.Sprintf("/%s/healthz", basePath), healthcheck.New())
    router.SetupRoutes(app)

    var wg sync.WaitGroup
    for _, port := range lsport {
        wg.Add(1)
        go func(p string) {
            defer wg.Done()
            if err := app.Listen(fmt.Sprintf(":%s", p), fiber.ListenConfig{
                EnablePrefork: true,
            }); err != nil {
                utils.Logger.Fatal().Err(err).Msg(fmt.Sprintf("Error starting server on port %s: %v", p, err))
            }
        }(port)
    }

    wg.Wait()
}

func initSwagger(app *fiber.App, basePath string) {
    b, _ := os.ReadFile("./public/css/theme-flattop.css")
    css := string(b)

    cfg := swaggo.Config{
        URL:          "doc.json",
        DeepLinking:  true,
        DocExpansion: "list",
        Title:        "Swagger E-Dashboard API",
        SyntaxHighlight: &swaggo.SyntaxHighlightConfig{
            Activate: true,
            Theme:    "arta",
        },
        CustomStyle:          template.CSS(css),
        PersistAuthorization: true,
    }

    app.Get(fmt.Sprintf("/%s/docs/*", basePath), swaggo.New(cfg))
    app.Get(fmt.Sprintf("/%s/static*", basePath), static.New("./public"))
}
