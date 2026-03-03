package lab

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"
    "fmt"
    "strings"
    "time"

    "github.com/gofiber/fiber/v3"
    "github.com/patrickmn/go-cache"
)

var ch *cache.Cache = cache.New(cache.NoExpiration, 30 * time.Minute)

func getFromCache(key string) []fiber.Map {
    var ls []fiber.Map
    if x, found := ch.Get(key); found {
        ls = x.([]fiber.Map)
    }

    return ls
}

func storeCache(key string, ls []fiber.Map) {
    if len(ls) > 0 {
        ch.Set(key, ls, cache.DefaultExpiration)
    } else {
        ch.Delete(key)
    }
}

// GetLabRadServiceClassReport
//
// @Tags Lab
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sRptType             query      string  true  "sRptType"
// @Param        COMPANY_ID           query      string  true  "COMPANY_ID"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /lab/service-cls [get]
func GetLabRadServiceClassReport(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sRptType := c.Query("sRptType")
    COMPANY_ID := c.Query("COMPANY_ID")
    vx_environment := c.Query("vx_environment")

    rpttype := strings.ToUpper(sRptType)
    cacheLabRadServiceClass := fmt.Sprintf("%s_LabRadServiceClass_%s", vx_user_session_id, rpttype)
    ls := getFromCache(cacheLabRadServiceClass)
    if ls != nil {
        return c.JSON(ls)
    }

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        userSessionID sql.NamedArg = sql.Named("userSessionID", vx_user_session_id)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
        reporttype    sql.NamedArg = sql.Named("reporttype", rpttype)
    )
    rows, err := db.Query(sqlx.LABRAD_SERVICE_CLASS, userSessionID, username, reporttype)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls = database.GetDataList(rows)
    if len(ls) < 1 {
        b := database.GenerateLabRadServiceClassData(db, sRptType, vx_user_session_id, vx_user_name, COMPANY_ID)
        if b {
            rows, err := db.Query(sqlx.LABRAD_SERVICE_CLASS, userSessionID, username, reporttype)
            if err != nil {
                utils.LogError(err)
                return err
            }

            defer rows.Close()

            ls = database.GetDataList(rows)
        } else {
            return fmt.Errorf("generateLabRadServiceClassData failed")
        }
    }

    storeCache(cacheLabRadServiceClass, ls)
    return c.JSON(ls)
}

// GetLabRadChargesReport
//
// @Tags Lab
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sRptType             query      string  true  "sRptType"
// @Param        COMPANY_ID           query      string  true  "COMPANY_ID"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /lab/charges [get]
func GetLabRadChargesReport(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sRptType := c.Query("sRptType")
    COMPANY_ID := c.Query("COMPANY_ID")
    vx_environment := c.Query("vx_environment")

    rpttype := strings.ToUpper(sRptType)
    cacheLabRadCharges := fmt.Sprintf("%s_LabRadCharges_%s", vx_user_session_id, rpttype)
    ls := getFromCache(cacheLabRadCharges)
    if ls != nil {
        return c.JSON(ls)
    }

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        userSessionID sql.NamedArg = sql.Named("userSessionID", vx_user_session_id)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
        reporttype    sql.NamedArg = sql.Named("reporttype", rpttype)
    )
    rows, err := db.Query(sqlx.LABRAD_CHARGES, userSessionID, username, reporttype)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls = database.GetDataList(rows)
    if len(ls) < 1 {
        b := database.GenerateLabRadChargesData(db, sRptType, vx_user_session_id, vx_user_name, COMPANY_ID)
        if b {
            rows, err := db.Query(sqlx.LABRAD_CHARGES, userSessionID, username, reporttype)
            if err != nil {
                utils.LogError(err)
                return err
            }

            defer rows.Close()

            ls = database.GetDataList(rows)
        } else {
            return fmt.Errorf("generateLabRadChargesData failed")
        }
    }

    storeCache(cacheLabRadCharges, ls)
    return c.JSON(ls)
}

// GetLabRadKPIStatisticsReport
//
// @Tags Lab
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sRptType             query      string  true  "sRptType"
// @Param        COMPANY_ID           query      string  true  "COMPANY_ID"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /lab/kpi-statistics [get]
func GetLabRadKPIStatisticsReport(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sRptType := c.Query("sRptType")
    COMPANY_ID := c.Query("COMPANY_ID")
    vx_environment := c.Query("vx_environment")

    rpttype := strings.ToUpper(sRptType)
    cacheLabRadKPIStatistics := fmt.Sprintf("%s_LabRadKPIStat_%s", vx_user_session_id, rpttype)
    ls := getFromCache(cacheLabRadKPIStatistics)
    if ls != nil {
        return c.JSON(ls)
    }

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    q := sqlx.LAB_KPI_STATISTICS
    if rpttype != "LABS" {
        q = sqlx.RADIOLOGY_KPI_STATISTICS
    }

    var (
        userSessionID sql.NamedArg = sql.Named("userSessionID", vx_user_session_id)
        username      sql.NamedArg = sql.Named("username", vx_user_name)
    )
    rows, err := db.Query(q, userSessionID, username)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls = database.GetDataList(rows)
    if len(ls) < 1 {
        b := false
        if rpttype == "LABS" {
            b = database.GenerateLabKPIStatisticsData(db, vx_user_session_id, vx_user_name, COMPANY_ID)
        } else {
            b = database.GenerateRadiologyKPIStatisticsData(db, vx_user_session_id, vx_user_name, COMPANY_ID)
        }

        if b {
            rows, err := db.Query(q, userSessionID, username)
            if err != nil {
                utils.LogError(err)
                return err
            }

            defer rows.Close()

            ls = database.GetDataList(rows)
        } else {
            return fmt.Errorf("generateKPIStatisticsData failed")
        }
    }

    storeCache(cacheLabRadKPIStatistics, ls)
    return c.JSON(ls)
}