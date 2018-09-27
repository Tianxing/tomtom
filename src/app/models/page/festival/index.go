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
    
    user := dataUser.User{}
    userInfo, err := user.GetInfoByUid(uid, stage)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%+v", userInfo), nil
}
