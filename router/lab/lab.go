package lab

import (
    "edashboard/controller/lab"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/lab")
    api.Get("/service-cls", lab.GetLabRadServiceClassReport)
    api.Get("/charges", lab.GetLabRadChargesReport)
    api.Get("/kpi-statistics", lab.GetLabRadKPIStatisticsReport)
}