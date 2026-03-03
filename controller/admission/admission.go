package admission

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

// GetAdmissionReportPopup
//
// @Tags Admission
// @Produce json
// @Param        sDate                query      string  true  "sDate"
// @Param        dateType             query      string  true  "dateType"
// @Param        sType                query      string  true  "sType"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Success 200
// @Router /admission/report-popup [get]
func GetAdmissionReportPopup(c fiber.Ctx) error {
    sDate := c.Query("sDate")
    dateType := c.Query("dateType")
    sType := c.Query("sType")
    vx_environment := c.Query("vx_environment")
    vx_user_session_id := c.Query("vx_user_session_id")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := ""
    switch dateType {
    case "7D", "30D":
        if sType == "WardPopup" {
            q = sqlx.GET_IP_ADMISSION_BYWARD_BYDAY
        }

        if sType == "DoctorPopup" {
            q = sqlx.GET_IP_ADMISSION_BYDOCTOR_BYDAY
        }
    case "6M", "12M":
        if sType == "WardPopup" {
            q = sqlx.GET_IP_ADMISSION_BYWARD_BYMONTH
        }

        if sType == "DoctorPopup" {
            q = sqlx.GET_IP_ADMISSION_BYDOCTOR_BYMONTH
        }
    }

    var (
        psDate          sql.NamedArg = sql.Named("sDate", sDate)
        user_session_id sql.NamedArg = sql.Named("user_session_id", vx_user_session_id)
    )
    rows, err := db.Query(q, psDate, user_session_id)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetAdmissionReportByMonth
//
// @Tags Admission
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        months               query      string  true  "months"
// @Param        sHospitalCode        query      string  true  "sHospitalCode"
// @Success 200
// @Router /admission/report/month [get]
func GetAdmissionReportByMonth(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    vx_environment := c.Query("vx_environment")
    months := c.Query("months")
    sHospitalCode := c.Query("sHospitalCode")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        user_name       sql.NamedArg = sql.Named("user_name", vx_user_name)
        user_session_id sql.NamedArg = sql.Named("user_session_id", vx_user_session_id)
        hospital_code   sql.NamedArg = sql.Named("hospital_code", sHospitalCode)
        month           sql.NamedArg = sql.Named("month", months)
    )
    res, err := db.Exec(sqlx.IP_ADMISSION_PRC_BY_MONTH, user_name, user_session_id, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    if res != nil {
        rows, err := db.Query(sqlx.GET_IP_ADMISSION_BY_MONTH, user_session_id, month)
        if err != nil {
            utils.LogError(err)
            return err
        }

        defer rows.Close()

        ls := database.GetDataList(rows)
        return c.JSON(ls)
    }

    return c.JSON([]any{})
}

// GetAdmissionReportByDay
//
// @Tags Admission
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        days                 query      string  true  "days"
// @Param        sHospitalCode        query      string  true  "sHospitalCode"
// @Success 200
// @Router /admission/report/day [get]
func GetAdmissionReportByDay(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    vx_environment := c.Query("vx_environment")
    days := c.Query("days")
    sHospitalCode := c.Query("sHospitalCode")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := ""
    switch days {
    case "7D":
        q = sqlx.IP_ADMISSION_PRC_BY_7DAY
    case "30D":
        q = sqlx.IP_ADMISSION_PRC_BY_30DAY
    }

    var (
        user_name       sql.NamedArg = sql.Named("user_name", vx_user_name)
        user_session_id sql.NamedArg = sql.Named("user_session_id", vx_user_session_id)
        hospital_code   sql.NamedArg = sql.Named("hospital_code", sHospitalCode)
    )
    res, err := db.Exec(q, user_name, user_session_id, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    if res != nil {
        rows, err := db.Query(sqlx.GET_IP_ADMISSION_BY_DAY, user_session_id)
        if err != nil {
            utils.LogError(err)
            return err
        }

        defer rows.Close()

        ls := database.GetDataList(rows)
        return c.JSON(ls)
    }

    return c.JSON([]any{})
}