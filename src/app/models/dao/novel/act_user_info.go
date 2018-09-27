package novel 

import (
    "fmt"
)

type ActUserInfo struct {
    Id int64
    Uid int64
    Cuid string
    Act_id int32
    Stage int64
    Gifts_count int64
    Platform int64
    Create_time int64
    Update_time int64
}

func (n *novel)GetActUserInfo(id int64) (ActUserInfo, error) {
    sql := "SELECT "
    sql += "`id`, `uid`, `cuid`, `act_id`, `stage`, "
    sql += "`gifts_count`, `platform`, "
    sql += "`create_time`, `update_time` "
    sql += "FROM "
    sql += TableInfo + " "
    sql += "WHERE "
    sql += "`id` = ?"

    info := ActUserInfo{}
    err:= n.QueryRow(sql, id).Scan(&info.Id, &info.Uid, &info.Cuid, &info.Act_id, &info.Stage, &info.Gifts_count, &info.Platform, &info.Create_time, &info.Update_time)

    return info, err 
}

func (n *novel)GetActUserInfoByUid(uid int64, stage int64) (ActUserInfo, error) {
    condition := make([]interface{}, 1, 2)
   
    sql := "SELECT "
    sql += "`id`, `uid`, `cuid`, `act_id`, `stage`, "
    sql += "`gifts_count`, `platform`, "
    sql += "`create_time`, `update_time` "
    sql += "FROM "
    sql += TableInfo + " "
    sql += "WHERE "
    sql += "`uid` = ? "

    condition[0] = uid

    if stage != 0 {
        sql += "AND `stage` = ?"
        condition = append(condition, stage)
    }

    fmt.Println(sql, uid, condition)
    info := ActUserInfo{}
    err:= n.QueryRow(sql, condition...).Scan(&info.Id, &info.Uid, &info.Cuid, &info.Act_id, &info.Stage, &info.Gifts_count, &info.Platform, &info.Create_time, &info.Update_time)
    
    return info, err
}
