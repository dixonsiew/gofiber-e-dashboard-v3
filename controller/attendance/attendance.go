package attendance

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

// GetOutpatientPopup
//
// @Tags Attendance
// @Produce json
// @Param        sDate                query      string  true  "sDate"
// @Param        dateType             query      string  true  "dateType"
// @Param        sType                query      string  true  "sType"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Success 200
// @Router /attendance/outpatient-popup [get]
func GetOutpatientPopup(c fiber.Ctx) error {
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
        if sType == "DoctorPopup" {
            q = sqlx.GET_OP_ATTENDANCE_BYDOCTOR_BYDAY
        }

        if sType == "ClinicPopup" {
            q = sqlx.GET_OP_ATTENDANCE_BYCLINIC_BYDAY
        }
    case "6M", "12M":
        if sType == "DoctorPopup" {
            q = sqlx.GET_OP_ATTENDANCE_BYDOCTOR_BYMONTH
        }

        if sType == "ClinicPopup" {
            q = sqlx.GET_OP_ATTENDANCE_BYCLINIC_BYMONTH
        }
    }

    rows, err := db.Query(q, sql.Named("sDate", sDate), sql.Named("user_session_id", vx_user_session_id))
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetOutpatientAttendancesByMonth
//
// @Tags Attendance
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        months               query      string  true  "months"
// @Param        sHospitalCode        query      string  true  "sHospitalCode"
// @Success 200
// @Router /attendance/outpatient/month [get]
func GetOutpatientAttendancesByMonth(c fiber.Ctx) error {
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
    res, err := db.Exec(sqlx.OP_ATTENDANCE_PRC_BY_MONTH, user_name, user_session_id, hospital_code)
    if err != nil {
        utils.LogError(err)
        return err
    }

    if res != nil {
        rows, err := db.Query(sqlx.GET_OP_ATTENDANCE_BY_MONTH, user_session_id, month)
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

// GetOutpatientAttendancesByDay
//
// @Tags Attendance
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Param        days                 query      string  true  "days"
// @Param        sHospitalCode        query      string  true  "sHospitalCode"
// @Success 200
// @Router /attendance/outpatient/day [get]
func GetOutpatientAttendancesByDay(c fiber.Ctx) error {
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
        q = sqlx.OP_ATTENDANCE_PRC_BY_7DAY
    case "30D":
        q = sqlx.OP_ATTENDANCE_PRC_BY_30DAY
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
        rows, err := db.Query(sqlx.GET_OP_ATTENDANCE_BY_DAY, user_session_id)
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