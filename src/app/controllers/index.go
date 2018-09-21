package controllers

import (
    "fmt"
    "net/http"
    "ff/controller"
)

type indexController struct {}

func init() {
    controller.Register("index", indexController{})
}

func (i *indexController)IndexAction(resp http.ResponseWriter, r *http.Request) {
//    fmt.Fprint(resp, r)
    fmt.Fprint(resp, "indexAction")
    return
}
