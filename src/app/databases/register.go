package databases

import (
    "database/sql"
    "errors"
)

var Conns = make(map[string]*sql.DB)
var ErrConnExisted = errors.New("DATABASE CONN HAS REGISTERED")

func Register(name string, conn *sql.DB) error {
    if _, exists := Conns[name]; exists {
        return ErrConnExisted
    }

    Conns[name] = conn
    return nil
}

func Close() {
    for _, conn := range Conns {
        conn.Close()
    }
}
