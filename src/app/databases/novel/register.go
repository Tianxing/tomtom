package novel

import (
    "fmt"
    "app/databases"
    "database/sql"
)

const (
    USERNAME = "root"
    PASSWORD = "mickey123"
    NETWORK = "tcp"
    HOST = "127.0.0.1"
    PORT = "3306"
    CHARSET = "utf8mb4"
    DATABASE = "novel"
)

func init() {

    dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, NETWORK, HOST, PORT, DATABASE, CHARSET)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("database novel register fail:" + err.Error())
        return 
    }

    databases.Register("novel", db)
    return
}
