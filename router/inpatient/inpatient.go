package inpatient

import (
    "edashboard/controller/inpatient"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/inpatient")
    api.Get("/bed-status", inpatient.GetBedStatus)
    api.Get("/bed-status-popup", inpatient.GetBedStatusPopup)
    api.Get("/specialty", inpatient.GetSpecialty)
    api.Get("/average-los", inpatient.GetAverageLOS)
    api.Get("/demographics", inpatient.GetDemographics)
    api.Get("/ward-census", inpatient.GetWardCensus)
    api.Get("/regadm", inpatient.GetRegistrationAndAdmission)
}