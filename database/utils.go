package database

import (
    "database/sql"
    "edashboard/utils"

    "github.com/gofiber/fiber/v3"
)

func GetDataList(rows *sql.Rows) []fiber.Map {
    cols, _ := rows.Columns()
    colTypes, _ := rows.ColumnTypes()
    lx := make([]fiber.Map, 0)

    values := make([]any, len(cols))
    for i, colType := range colTypes {
        var val any
        switch colType.DatabaseTypeName() {
        case "NCHAR", "VARCHAR", "VARCHAR2", "NVARCHAR2":
            val = new(sql.NullString)
        case "INTEGER", "LONG", "SMALLINT":
            val = new(sql.NullInt64)
        case "REAL", "DOUBLE PRECISION", "FLOAT", "DECIMAL", "NUMBER":
            val = new(sql.NullFloat64)
        default:
            val = new(interface{})
        }
        values[i] = val
    }

    for rows.Next() {
        err := rows.Scan(values...)
        if err != nil {
            utils.LogError(err)
        }

        mx := make(fiber.Map)
        for i, col := range cols {
            switch v := values[i].(type) {
            case *sql.NullString:
                mx[col] = v.String
            case *sql.NullInt64:
                mx[col] = v.Int64
            case *sql.NullFloat64:
                mx[col] = v.Float64
            }
        }
        lx = append(lx, mx)
    }

    return lx
}