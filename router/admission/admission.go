package admission

import (
    "edashboard/controller/admission"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/admission")
    api.Get("/report-popup", admission.GetAdmissionReportPopup)
    api.Get("/report/month", admission.GetAdmissionReportByMonth)
    api.Get("/report/day", admission.GetAdmissionReportByDay)
}