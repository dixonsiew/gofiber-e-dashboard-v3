package finance

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"
    "fmt"
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

// GetARAgeingReport
//
// @Tags Finance
// @Produce json
// @Param        vx_user_session_id   query      string  true  "vx_user_session_id"
// @Param        vx_user_name         query      string  true  "vx_user_name"
// @Param        sType                query      string  true  "sType"
// @Param        sTopARNum            query      string  true  "sTopARNum"
// @Param        COMPANY_ID           query      string  true  "COMPANY_ID"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /finance/ar-ageing [get]
func GetARAgeingReport(c fiber.Ctx) error {
    vx_user_session_id := c.Query("vx_user_session_id")
    vx_user_name := c.Query("vx_user_name")
    sType := c.Query("sType")
    sTopARNum := c.Query("sTopARNum")
    COMPANY_ID := c.Query("COMPANY_ID")
    vx_environment := c.Query("vx_environment")

    q := sqlx.AR_AGEING_OVERALL
    switch sType {
    case "P", "I":
        q = sqlx.AR_AGEING
    case "T":
        q = sqlx.AR_AGEING_TOP
    }

    customertype := sType
    t := sType
    switch sType {
    case "P", "I":
        customertype = sType
    case "T":
        t = "I"
        customertype = t
    }

    var total float64 = 0
    key := fmt.Sprintf("%s_FinanceOverall", vx_user_session_id)

    switch sType {
    case "P":
        key = fmt.Sprintf("%s_FinancePatient", vx_user_session_id)
    case "I":
        key = fmt.Sprintf("%s_Finance3rdParty", vx_user_session_id)
    case "T":
        key = fmt.Sprintf("%s_FinanceTop_%s", vx_user_session_id, sTopARNum)
    }

    ls := getFromCache(key)
    if ls != nil {
        if sType == "T" {
            for _, m := range ls {
                i := m["AMT_TOTAL"].(float64)
                total += i
            }
            return c.JSON(fiber.Map{
                "data":  ls,
                "total": total,
            })
        } else {
            if hasTotal(ls) {
                return c.JSON(ls)
            } else {
                ls = make([]fiber.Map, 0)
                storeCache(key, ls)
            }
        }
    }

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        usersessionid sql.NamedArg = sql.Named("usersessionid", vx_user_session_id)
        userid        sql.NamedArg = sql.Named("userid", vx_user_name)
        pTopARCount   sql.NamedArg = sql.Named("TopARCount", sTopARNum)
        pcustomertype sql.NamedArg = sql.Named("customertype", customertype)
    )
    params := []any{ usersessionid, userid, pTopARCount, pcustomertype }
    rows, err := db.Query(q, params...)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()
    ls = database.GetDataList(rows)

    if !hasTotal(ls) {
        bP := database.GenerateARAgeingReportData(db, "P", vx_user_session_id, vx_user_name, COMPANY_ID)
        bI := database.GenerateARAgeingReportData(db, "I", vx_user_session_id, vx_user_name, COMPANY_ID)
        if bP && bI {
            rows, err := db.Query(q, params...)
            if err != nil {
                utils.LogError(err)
                return err
            }

            defer rows.Close()
            ls = database.GetDataList(rows)
        } else {
            return fmt.Errorf("generateARAgeingReportData failed")
        }
    } else {
        storeCache(key, ls)
    }

    if sType == "T" {
        total = 0
        for _, m := range ls {
            i := m["AMT_TOTAL"].(float64)
            total += i
        }
        return c.JSON(fiber.Map{
            "data":  ls,
            "total": total,
        })
    }

    return c.JSON(ls)
}

func hasTotal(ls []fiber.Map) bool {
    b := false
    if len(ls) > 0 {
        m := ls[0]
        _, ok := m["Total"]
        if ok {
            x := m["Total"].(float64)
            if x > 0 {
                b = true
            }
        }
    }
    return b
}