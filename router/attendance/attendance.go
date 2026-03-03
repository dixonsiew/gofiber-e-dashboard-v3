package attendance

import (
    "edashboard/controller/attendance"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/attendance")
    api.Get("/outpatient-popup", attendance.GetOutpatientPopup)
    api.Get("/outpatient/month", attendance.GetOutpatientAttendancesByMonth)
    api.Get("/outpatient/day", attendance.GetOutpatientAttendancesByDay)
}