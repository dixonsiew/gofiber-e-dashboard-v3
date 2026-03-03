package referral

import (
    "edashboard/controller/referral"
    "github.com/gofiber/fiber/v3"
)

func SetupRoutes(router fiber.Router) {
    api := router.Group("/referral")
    api.Get("/external", referral.GetExternalReport)
    api.Get("/external-popup", referral.GetExternalPopup)
    api.Get("/internal", referral.GetInternalReport)
    api.Get("/internal-popup", referral.GetInternalPopup)
}