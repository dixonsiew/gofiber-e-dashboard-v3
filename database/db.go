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

var dbMap map[string]*sql.DB = make(map[string]*sql.DB)

func InitDB() {
    s := config.Config("pools")
    ls := strings.SplitSeq(s, ",")
    for k := range ls {
        ConnectDB(k)
    }
}

func SetDb(db *sql.DB, s string) {
    dbMap[s] = db
}

func GetDb(s string) *sql.DB {
    db, ok := dbMap[s]
    if !ok {
        ConnectDB(s)
    }

    return db
}

func ConnectDB(s string) {
    username := config.Config(fmt.Sprintf("%s.db.username", s))
    pwd := config.Config(fmt.Sprintf("%s.db.pwd", s))
    url := config.Config(fmt.Sprintf("%s.db.url", s))
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
        SetDb(db, s)
        utils.LogInfo("Connection Opened to Database")
    }
}

func CloseDB() {
    for _, v := range dbMap {
        v.Close()
    }
}