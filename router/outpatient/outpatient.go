package outpatient

import (
    "edashboard/controller/outpatient"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/outpatient")
    api.Get("/clinic", outpatient.GetClinics)
    api.Get("/clinic-popup", outpatient.GetClinicPopup)
    api.Get("/attendances", outpatient.GetAttendances)
    api.Get("/demographics", outpatient.GetDemographics)
}