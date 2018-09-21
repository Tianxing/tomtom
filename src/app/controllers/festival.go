package controllers

import(
    "fmt"
    "net/http"
    "ff/controller"
    "utils/festival"
)

type festivalController struct{}

func init() {
    controller.Register("festival", festivalController{})
}

func (f *festivalController)IndexAction(resp http.ResponseWriter, r *http.Request) {
    stageNow := festival.Stage()
    fmt.Fprintln(resp, stageNow)
}
