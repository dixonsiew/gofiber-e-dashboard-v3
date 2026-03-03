package mortality

import (
    "database/sql"
    "edashboard/database"
    sqlx "edashboard/sql"
    "edashboard/utils"
    "strings"

    "github.com/gofiber/fiber/v3"
)

// GetTotalDeath
//
// @Tags Mortality
// @Produce json
// @Param        sYear                query      string  true  "sYear"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /mortality/total-death [get]
func GetTotalDeath(c fiber.Ctx) error {
    sYear := c.Query("sYear")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        year sql.NamedArg = sql.Named("year", sYear)
    )
    rows, err := db.Query(sqlx.MORTALITY_TOTAL_DEATH, year)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetTotalDeathGender
//
// @Tags Mortality
// @Produce json
// @Param        sYear                query      string  true  "sYear"
// @Param        sGender              query      string  true  "sGender"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /mortality/total-death-gender [get]
func GetTotalDeathGender(c fiber.Ctx) error {
    sYear := c.Query("sYear")
    sGender := c.Query("sGender")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        year   sql.NamedArg = sql.Named("year", sYear)
        gender sql.NamedArg = sql.Named("gender", strings.ToUpper(sGender))
    )
    rows, err := db.Query(sqlx.MORTALITY_TOTAL_DEATH_BY_GENDER, year, gender)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetTotalDeathCountGender
//
// @Tags Mortality
// @Produce json
// @Param        sYear                query      string  true  "sYear"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /mortality/total-death-count-gender [get]
func GetTotalDeathCountGender(c fiber.Ctx) error {
    sYear := c.Query("sYear")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        year sql.NamedArg = sql.Named("year", sYear)
    )
    rows, err := db.Query(sqlx.MORTALITY_TOTAL_DEATH_COUNT_BY_GENDER, year)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetTotalDeathCase
//
// @Tags Mortality
// @Produce json
// @Param        sYear                query      string  true  "sYear"
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /mortality/total-death-case [get]
func GetTotalDeathCase(c fiber.Ctx) error {
    sYear := c.Query("sYear")
    vx_environment := c.Query("vx_environment")

    db := database.GetDbrs(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    var (
        year sql.NamedArg = sql.Named("year", sYear)
    )
    rows, err := db.Query(sqlx.MORTALITY_TOTAL_DEATH_CASE, year)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}

// GetGender
//
// @Tags Mortality
// @Produce json
// @Param        vx_environment       query      string  true  "vx_environment"
// @Success 200
// @Router /mortality/lookup/gender [get]
func GetGender(c fiber.Ctx) error {
    vx_environment := c.Query("vx_environment")

    db := database.GetDb(vx_environment)
    if db == nil {
        utils.LogInfo("db is nil")
    }

    rows, err := db.Query(`SELECT VALUE, MEANING FROM NH_DSHB_LOOKUP_CODE WHERE LOOKUP_TYPE = 'SEX'`)
    if err != nil {
        utils.LogError(err)
        return err
    }

    defer rows.Close()

    ls := database.GetDataList(rows)
    return c.JSON(ls)
}