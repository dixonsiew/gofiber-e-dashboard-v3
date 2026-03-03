package outpatient

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

// GetClinics
//
// @Tags Outpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /outpatient/clinic [get]
func GetClinics(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sType := c.Query("sType")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.OP_CLINIC_QUEUE
    if sType == "OPConsultant" {
        q = sqlx.OP_CONSULTANT
    }

    stoken := database.GetToken(db, vx_user_session_id, vx_user_name)
    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
        token         sql.NamedArg = sql.Named("token", stoken)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
    )
    rows, err := db.Query(q, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetClinicPopup
//
// @Tags Outpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        sClinicCode          query      string  true  "sClinicCode"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /outpatient/clinic-popup [get]
func GetClinicPopup(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sType := c.Query("sType")
    sClinicCode := c.Query("sClinicCode")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.OP_CLINIC_QUEUE_POPUP
    if sType == "OPConsultantPopup" {
        q = sqlx.OP_CONSULTANT_POPUP
    }

    stoken := database.GetToken(db, vx_user_session_id, vx_user_name)
    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
        token         sql.NamedArg = sql.Named("token", stoken)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
        clinicCode    sql.NamedArg = sql.Named("clinicCode", sClinicCode)
    )
    rows, err := db.Query(q, hospital_code, token, username, clinicCode)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetAttendances
//
// @Tags Outpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /outpatient/attendances [get]
func GetAttendances(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    stoken := database.GetToken(db, vx_user_session_id, vx_user_name)
    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
        token         sql.NamedArg = sql.Named("token", stoken)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
    )
    rows, err := db.Query(sqlx.OP_ATTENDANCES, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetDemographics
//
// @Tags Outpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /outpatient/demographics [get]
func GetDemographics(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sType := c.Query("sType")
    vx_comp_name := c.Query("vx_comp_name")
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.OP_DEMOGRAPHICS_NATIONALITY
    switch sType {
    case "PaymentClass":
        q = sqlx.OP_DEMOGRAPHICS_PAYMENTCLASS
    case "PatientType":
        q = sqlx.OP_DEMOGRAPHICS_PATIENTTYPE
    }

    stoken := database.GetToken(db, vx_user_session_id, vx_user_name)
    var (
        hospital_code sql.NamedArg = sql.Named("hospital_code", vx_comp_name)
        token         sql.NamedArg = sql.Named("token", stoken)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
    )
    rows, err := db.Query(q, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}