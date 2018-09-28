package festival

import (
    "fmt"
    utils "utils/festival"
    "passport"
    dataUser "app/models/data/user"
//    "strconv"
)

type Index struct{}

func (i *Index)Exec(uss passport.Session) (string, error) {
    
    stage := utils.Stage()
    uid := uss.Uid()
    
    user := dataUser.New()
    userInfo, err := user.GetActUserInfoByUid(uid, stage)
    gifts, err := user.GetActUserDrawgiftInfoByUid(uid, stage)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%+v\n\n%+v", userInfo, gifts), nil
}
