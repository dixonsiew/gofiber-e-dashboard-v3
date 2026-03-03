package database

import (
	"database/sql"
	"edashboard/utils"

	"github.com/sijms/go-ora/v2"
)

func GetToken(con *sql.DB, sessionid string, username string) string {
    stoken := ""
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        token          go_ora.NVarChar
    )
    q := `BEGIN NH_GET_DASHBOARD_VIEW_TOKEN_SP(:gUserSessionId, :gUserName, :token); END;`
    _, err := con.Exec(q, gUserSessionId, gUserName, go_ora.Out{Dest: &token, Size: 5000})
    if err != nil {
        utils.LogError(err)
        return stoken
    }

    stoken = string(token)
    return stoken
}

func IsWardCensusReportDataExist(con *sql.DB, sessionid string, username string) int {
    n := 0
    var (
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
    )
    q := `SELECT COUNT(*) AS TOTALRECORDS FROM NH_RPT_WARD_CENSUS_TEMP WHERE USER_ID = :gUserName and USER_SESSION_ID = :gUserSessionId`
    rows, err := con.Query(q, gUserName, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return 0
    }

    defer rows.Close()

    if rows.Next() {
        var i sql.NullInt32
        err := rows.Scan(&i)
        if err != nil {
            utils.LogError(err)
        }

        n = int(i.Int32)
    }

    return n
}

func GenerateWardCensusReportData(con *sql.DB, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
    )
    q := `BEGIN NH_RPT_WARD_CENSUS_SP(SYSDATE, :gUserName, :gUserSessionId, :gCompanyId); END;`
    res, err := con.Exec(q, gUserName, gUserSessionId, gCompanyId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}

func GenerateARAgeingReportData(con *sql.DB, custtype string, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
        gCustomerType  sql.NamedArg = sql.Named("gCustomerType", custtype)
        gReportType    sql.NamedArg = sql.Named("gReportType", "S")
    )
    q := `BEGIN NH_RPT_AR_AGING_SP(:gCustomerType, :gReportType, NULL, NULL, NULL, :gCompanyId, :gUserName, :gUserSessionId); END;`
    res, err := con.Exec(q, gCustomerType, gReportType, gCompanyId, gUserName, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}

func GenerateLabRadServiceClassData(con *sql.DB, reporttype string, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
        gReportType    sql.NamedArg = sql.Named("gReportType", reporttype)
    )
    q := `BEGIN NH_RPT_SERVICE_STATISTICS(:gUserName, :gReportType, :gCompanyId, :gUserSessionId, TO_DATE(SYSDATE, 'DD-MON-YY'), TO_DATE(SYSDATE, 'DD-MON-YY')); END;`
    res, err := con.Exec(q, gUserName, gReportType, gCompanyId, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}

func GenerateLabRadChargesData(con *sql.DB, reporttype string, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
        gReportType    sql.NamedArg = sql.Named("gReportType", reporttype)
    )
    q := `BEGIN NH_RPT_SERVICE_CHARGES(:gUserName, :gReportType, :gCompanyId, :gUserSessionId, TO_DATE(SYSDATE, 'DD-MON-YY'), TO_DATE(SYSDATE, 'DD-MON-YY')); END;`
    res, err := con.Exec(q, gUserName, gReportType, gCompanyId, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}

func GenerateLabKPIStatisticsData(con *sql.DB, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
    )
    q := `BEGIN NH_RPT_LAB_KPI_STATS(:gUserName, :gCompanyId, :gUserSessionId, TO_DATE(SYSDATE, 'DD-MON-YY'), TO_DATE(SYSDATE, 'DD-MON-YY'), NULL, NULL); END;`
    res, err := con.Exec(q, gUserName, gCompanyId, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}

func GenerateRadiologyKPIStatisticsData(con *sql.DB, sessionid string, username string, companyid string) bool {
    var (
        gUserSessionId sql.NamedArg = sql.Named("gUserSessionId", sessionid)
        gUserName      sql.NamedArg = sql.Named("gUserName", username)
        gCompanyId     sql.NamedArg = sql.Named("gCompanyId", companyid)
    )
    q := `BEGIN NH_RPT_RADIOLOGY_KPI_STATS(:gUserName, :gCompanyId, :gUserSessionId, TO_DATE(SYSDATE, 'DD-MON-YY'), TO_DATE(SYSDATE, 'DD-MON-YY'), NULL, NULL); END;`
    res, err := con.Exec(q, gUserName, gCompanyId, gUserSessionId)
    if err != nil {
        utils.LogError(err)
        return false
    }

    if res != nil {
        return true
    }

    return false
}