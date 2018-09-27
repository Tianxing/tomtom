package controllers

import(
    "fmt"
    "net/http"
    "ff/controller"
    "passport"
    pages "app/models/page/festival" 
)

type festivalController struct{}

func init() {
    controller.Register("festival", festivalController{})
}

func (f *festivalController)IndexAction(resp http.ResponseWriter, r *http.Request) {
    uss, err  := passport.CheckLogin(r)

    if err != nil {
        uss, _ = passport.Login(resp)
    }

    pInfo, err := (&pages.Index{}).Exec(uss)
    fmt.Fprintln(resp, pInfo, err)
}
