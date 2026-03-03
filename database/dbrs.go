package database

import (
    "database/sql"
    "edashboard/config"
    "edashboard/utils"
    "fmt"
    "strings"
    "time"

    _ "github.com/sijms/go-ora/v2"
)

var dbrsMap map[string]*sql.DB = make(map[string]*sql.DB)

func InitDBRs() {
    s := config.Config("pools")
    ls := strings.Split(s, ",")
    for _, k := range ls {
        ConnectDBRs(k)
    }
}

func SetDbrs(db *sql.DB, s string) {
    dbrsMap[s] = db
}

func GetDbrs(s string) *sql.DB {
    db, ok := dbrsMap[s]
    if !ok {
        ConnectDBRs(s)
    }

    return db
}

func ConnectDBRs(s string) {
    h := fmt.Sprintf("%s.db.rs.host", s)
    if h == "" {
        return
    }
    
    username := config.Config(fmt.Sprintf("%s.db.rs.username", s))
    pwd := config.Config(fmt.Sprintf("%s.db.rs.pwd", s))
    url := config.Config(fmt.Sprintf("%s.db.rs.url", s))
    connStr := fmt.Sprintf("oracle://%s:%s@%s", username, pwd, url)
    db, err := sql.Open("oracle", connStr)
    // dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s" heterogeneousPool=false standaloneConnection=false`, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_URL"))
    // fmt.Println(dsn)
    // DB, err := sql.Open("godror", dsn)

    if err != nil {
        utils.LogError(err)
    } else {
        db.SetMaxOpenConns(50)
        db.SetMaxIdleConns(5)
        db.SetConnMaxLifetime(5 * time.Minute)
        SetDbrs(db, s)
        utils.LogInfo("Connection Opened to Database")
    }
}

func CloseDBRs() {
    for _, v := range dbrsMap {
        v.Close()
    }
}