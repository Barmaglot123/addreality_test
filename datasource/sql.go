package datasource

import (
    "database/sql"
    "fmt"
)

var DB *sql.DB

const (
    host     = "localhost"
    port     = 5432
    user     = "test_user"
    password = "your-rUeIZVWr"
    dbname   = "test_task"
)

func SetupSql() {
    var err error
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    DB, err = sql.Open("postgres", psqlInfo)

    if err != nil {
        panic(err)
    }
}