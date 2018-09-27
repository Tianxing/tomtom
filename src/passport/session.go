package passport

import (
//    "fmt"
    "net/http"
    "time"
    "errors"
    "strconv"
)

type Session interface{
    Uid() int64
    Name() string
}

type session struct {
    uid int64
    name string
}

func (u *session)Uid() int64 {
    return u.uid
}

func (u *session)Name() string {
    return u.name
}

var ErrNoLogin = errors.New("ERROR NO LOGIN")
var ErrFailLogin = errors.New("ERROR FAIL LOGIN")

func CheckLogin(r *http.Request) (Session, error) {
    uss, err := r.Cookie("uss")
    if err != nil {

        return nil, ErrNoLogin
    }
    uid, _ := strconv.ParseInt(uss.Value, 10, 64)

    return &session{uid, "tom"}, nil
}

func Login(resp http.ResponseWriter) (Session, error) {

    after1min := time.Now()
    min1, _ := time.ParseDuration("1min")
    after1min.Add(min1)
    ussCookie := http.Cookie{
        Name: "uss",
        Value: "12306",
        Path: "/",
        Domain: "tt.com",
        MaxAge: 300,
    }

    http.SetCookie(resp, &ussCookie)

    if !true {
        return nil, ErrFailLogin
    }
    
    return &session{12306, "tom"}, nil
}
