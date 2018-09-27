package user

import(
    dao "app/models/dao/novel"
)

type User struct{}

func (u *User)GetInfoByUid(uid int64, stage int64) (dao.ActUserInfo, error) {
    daoNovel := dao.New()
    return daoNovel.GetActUserInfoByUid(uid, stage)
}
