package user

import(
    dao "app/models/dao/novel"
)

type user struct{
    dao *dao.Novel
}

func New() *user {
    return &user{dao:dao.New()}
}

func (u *user)GetActUserInfoByUid(uid int64, stage int64) (dao.ActUserInfo, error) {

    return u.dao.GetActUserInfoByUid(uid, stage)
}

func (u *user)GetActUserDrawgiftInfoByUid(uid int64, stage int64)([]dao.ActUserDrawgiftInfo, error) {

    return u.dao.GetActUserDrawgiftInfoByUid(uid, stage)
}
