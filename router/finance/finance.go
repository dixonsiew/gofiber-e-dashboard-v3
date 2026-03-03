package finance

import (
    "edashboard/controller/finance"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/finance")
    api.Get("/ar-ageing", finance.GetARAgeingReport)
}