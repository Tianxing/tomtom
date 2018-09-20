package controller
import (
    net/http
)

function IndexAction(resp http.ResponseWriter, r *http.Request) {
    fmt.Fprint("indexAction")
}
