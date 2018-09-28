package novel

type ActUserDrawgiftInfo struct {
    Id int64
    Uid int64
    Act_id int
    Stage int64
    Gift_id int
    Gift_count int
    Create_time int
    Update_time int
}

func (n *Novel)GetActUserDrawgiftInfoByUid(uid int64, stage int64) ([]ActUserDrawgiftInfo, error) {
    condition := make([]interface{}, 1, 2)

    sql := "SELECT "
    sql += "`id`, `uid`, `act_id`, "
    sql += "`stage`, `gift_id`, `gift_count`, "
    sql += "`create_time`, `update_time` "
    sql += "FROM "
    sql += TableActUserDrawgiftInfo + " "
    sql += "WHERE "
    sql += "`uid` = ? "

    condition[0] = uid

    if stage != 0 {
        sql += "AND `stage` = ? "
        condition = append(condition, stage)
    }

    var result []ActUserDrawgiftInfo
    rows, err := n.Query(sql, condition...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var info = ActUserDrawgiftInfo{}
        err := rows.Scan(
          &info.Id, &info.Uid, &info.Act_id,
          &info.Stage, &info.Gift_id, &info.Gift_count,
          &info.Create_time, &info.Update_time,
        )
        if err != nil {
            return nil, err
        }
        result = append(result, info)
    }
    //close quickly
    rows.Close()
    return result, nil
}
