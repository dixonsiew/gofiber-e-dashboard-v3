package appointment

import (
    "edashboard/controller/appointment"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/appointment")
    api.Get("/statistics", appointment.GetAppointmentStatistics)
    api.Get("/statistics-popup", appointment.GetAppointmentStatisticsPopup)
}