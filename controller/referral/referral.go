package referral

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

// GetExternalReport
//
// @Tags Referral
// @Produce json
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /referral/external [get]
func GetExternalReport(c fiber.Ctx) error {
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
    )
    rows, err := db.Query(sqlx.REFERRAL_EXTERNAL, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetExternalPopup
//
// @Tags Referral
// @Produce json
// @Param        sDate                query      string  true  "sDate"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /referral/external-popup [get]
func GetExternalPopup(c fiber.Ctx) error {
    sDate := c.Query("sDate")
    sType := c.Query("sType")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.REFERRAL_EXTERNAL_BY_DOCTOR
    if sType == "SpecialtyPopup" {
        q = sqlx.REFERRAL_EXTERNAL_BY_SPECIALTY
    }

    var (
        dt            sql.NamedArg = sql.Named("dt", sDate)
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
    )
    rows, err := db.Query(q, dt, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetInternalReport
//
// @Tags Referral
// @Produce json
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /referral/internal [get]
func GetInternalReport(c fiber.Ctx) error {
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
    )
    rows, err := db.Query(sqlx.REFERRAL_INTERNAL, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetInternalPopup
//
// @Tags Referral
// @Produce json
// @Param        sDate                query      string  true  "sDate"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /referral/internal-popup [get]
func GetInternalPopup(c fiber.Ctx) error {
    sDate := c.Query("sDate")
    sType := c.Query("sType")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.REFERRAL_INTERNAL_BY_SENDER
    if sType == "RecvPopup" {
        q = sqlx.REFERRAL_INTERNAL_BY_RECEIVING
    }

    var (
        dt            sql.NamedArg = sql.Named("dt", sDate)
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
    )
    rows, err := db.Query(q, dt, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}