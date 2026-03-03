package mortality

import (
    "edashboard/controller/mortality"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/mortality")
    api.Get("/total-death", mortality.GetTotalDeath)
    api.Get("/total-death-gender", mortality.GetTotalDeathGender)
    api.Get("/total-death-count-gender", mortality.GetTotalDeathCountGender)
    api.Get("/total-death-case", mortality.GetTotalDeathCase)
    api.Get("/lookup/gender", mortality.GetGender)
}