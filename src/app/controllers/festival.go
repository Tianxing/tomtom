package controllers

import(
    "fmt"
    "net/http"
    "ff/controller"
    "utils/festival"
    "passport"
)

type festivalController struct{}

func init() {
    controller.Register("festival", festivalController{})
}

func (f *festivalController)IndexAction(resp http.ResponseWriter, r *http.Request) {
    stageNow := festival.Stage()
    uss, _ := passport.CheckLogin(resp, r)

    fmt.Fprintln(resp, "cookie;", r.Cookies())
    fmt.Fprintln(resp, "stage:", stageNow)
    fmt.Fprintln(resp, "uss:")
    fmt.Fprintln(resp, uss)
}
