package inpatient

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"
    "fmt"

    "github.com/gofiber/fiber/v3"
)

// GetBedStatus
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/bed-status [get]
func GetBedStatus(c fiber.Ctx) error {
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
    rows, err := db.Query(sqlx.IP_BED_STATUS, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetBedStatusPopup
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sBedStatus           query      string  true  "sBedStatus"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/bed-status-popup [get]
func GetBedStatusPopup(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sBedStatus := c.Query("sBedStatus")
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
        bed_status    sql.NamedArg = sql.Named("bed_status", sBedStatus)
    )
    rows, err := db.Query(sqlx.BED_STATUS_BY_WARD, hospital_code, token, username, bed_status)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetSpecialty
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/specialty [get]
func GetSpecialty(c fiber.Ctx) error {
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
    rows, err := db.Query(sqlx.IP_BY_SPECIALTY, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetAverageLOS
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/average-los [get]
func GetAverageLOS(c fiber.Ctx) error {
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
    rows, err := db.Query(sqlx.IP_AVG_LOS, hospital_code, token, username)
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
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/demographics [get]
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

    q := sqlx.IP_DEMOGRAPHICS_NATIONALITY
    switch sType {
    case "Gender":
        q = sqlx.IP_DEMOGRAPHICS_GENDER
    case "Age":
        q = sqlx.IP_DEMOGRAPHICS_AGE
    case "PaymentClass":
        q = sqlx.IP_DEMOGRAPHICS_PAYMENTCLASS
    case "Race":
        q = sqlx.IP_DEMOGRAPHICS_RACE
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

// GetWardCensus
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        COMPANY_ID           query      string  true  "COMPANY_ID"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/ward-census [get]
func GetWardCensus(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sType := c.Query("sType")
    COMPANY_ID := c.Query("COMPANY_ID")
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    ls := make([]fiber.Map, 0)
    n := database.IsWardCensusReportDataExist(db, vx_user_session_id, vx_user_name)
    if n > 0 {
        ls = getWardCensusReport(db, vx_user_session_id, vx_user_name, sType)
    } else {
        b := database.GenerateWardCensusReportData(db, vx_user_session_id, vx_user_name, COMPANY_ID)
        if b {
            ls = getWardCensusReport(db, vx_user_session_id, vx_user_name, sType)
        } else {
            return fmt.Errorf("generateWardCensusReportData failed")
        }
    }

    return c.JSON(ls)
}

// GetRegistrationAndAdmission
//
// @Tags Inpatient
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_comp_name         query      string  true  "vx_comp_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /inpatient/regadm [get]
func GetRegistrationAndAdmission(c fiber.Ctx) error {
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
    rows, err := db.Query(sqlx.REGISTRATION_ADMISSION, hospital_code, token, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

func getWardCensusReport(db *sql.DB, sUserSessionId string, sUsername string, sType string) []fiber.Map {
    q := sqlx.WARD_CENSUS_BY_WARD
    switch sType {
    case "WardCensusAge":
        q = sqlx.WARD_CENSUS_BY_AGE
    case "WardCensusPaymentClass":
        q = sqlx.WARD_CENSUS_BY_PAYMENTCLASS
    }

    ls := make([]fiber.Map, 0)
    var (
        usersessionid sql.NamedArg = sql.Named("usersessionid", sUserSessionId)
        userid        sql.NamedArg = sql.Named("userid", sUsername)
    )
    rows, err := db.Query(q, usersessionid, userid)
    if err != nil {
        utils.LogError(err)
        return ls
    }

    defer rows.Close()

    ls = database.GetDataList(rows)
    return ls
}