package novel

import (
    "dao"
    "app/databases"
)

const DbName = "novel"
const TableInfo = "novel_act_user_info"

type novel struct {
   *dao.Dao
}

func New() *novel {
    db, exists := databases.Conns[DbName]
    if !exists {
        return nil
    }

    return &novel{Dao:&dao.Dao{DB:db}}
}
