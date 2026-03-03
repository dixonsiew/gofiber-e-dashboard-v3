package router

import (
    "edashboard/router/admission"
    "edashboard/router/appointment"
    "edashboard/router/attendance"
    "edashboard/router/finance"
    "edashboard/router/inpatient"
    "edashboard/router/lab"
    "edashboard/router/mortality"
    "edashboard/router/outpatient"
    "edashboard/router/referral"
    "github.com/gofiber/fiber/v3"
    // "github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
    api := app.Group("/edashboard")
    admission.SetupRoutes(api)
    appointment.SetupRoutes(api)
    attendance.SetupRoutes(api)
    finance.SetupRoutes(api)
    inpatient.SetupRoutes(api)
    lab.SetupRoutes(api)
    mortality.SetupRoutes(api)
    outpatient.SetupRoutes(api)
    referral.SetupRoutes(api)
}