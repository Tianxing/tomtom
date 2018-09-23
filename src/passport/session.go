package passport

import (
    "fmt"
    "net/http"
    "time"
)

func CheckLogin(resp http.ResponseWriter, r *http.Request) (string, error) {
    uss, err := r.Cookie("uss")
    fmt.Println(err)
    if err == nil {
        fmt.Println("GET USS FROM COOKIE")
        return uss.String(), nil
    }

    after1min := time.Now()
    min1, _ := time.ParseDuration("1hour")
    after1min.Add(min1)
    ussCookie := http.Cookie{
        Name: "uss", 
        Value: "vv12306", 
        Path: "/", 
        Domain: "tt.com", 
        MaxAge: 300,
    }

    http.SetCookie(resp, &ussCookie)
    return "vv12306", nil
}
