package appointment

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

// GetAppointmentStatistics
//
// @Tags Appointment
// @Produce json
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /appointment/statistics [get]
func GetAppointmentStatistics(c fiber.Ctx) error {
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    rows, err := db.Query(sqlx.APPOINTMENT_STATISTICS, sql.Named("hospital_code", vx_comp_name))
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetAppointmentStatisticsPopup
//
// @Tags Appointment
// @Produce json
// @Param        sDate                query      string  true  "sDate"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /appointment/statistics-popup [get]
func GetAppointmentStatisticsPopup(c fiber.Ctx) error {
    sDate := c.Query("sDate")
    sType := c.Query("sType")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.APPOINTMENT_STATISTICS_BY_DOCTOR
    if sType == "SpecialtyPopup" {
        q = sqlx.APPOINTMENT_STATISTICS_BY_SPECIALTY
    }

    rows, err := db.Query(q, sql.Named("dt", sDate), sql.Named("hospital_code", vx_comp_name))
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}