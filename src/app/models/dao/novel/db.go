package novel

import (
    "dao"
    "app/databases"
)

const DbName = "novel"
const TableActUserInfo = "novel_act_user_info"
const TableActUserDrawgiftInfo = "novel_act_user_drawgift_info"

type Novel struct {
   *dao.Dao
}

func New() *Novel {
    db, exists := databases.Conns[DbName]
    if !exists {
        return nil
    }

    return &Novel{Dao:&dao.Dao{DB:db}}
}
