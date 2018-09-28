package main

import (
  "net/http"
  "net/url"
  "ff/controller"
  "log"
  _"app/controllers"
  "os"
  "app/databases"
  _"github.com/go-sql-driver/mysql"
  _"app/databases/novel"
)

const (
  APPNAME = "tomtom"
  LOG_PATH = "/Users/baitianxing/log/" + APPNAME + "/"
  LOG_FILE = LOG_PATH + APPNAME + ".log"
  LOG_FILE_WF = LOG_PATH + APPNAME + ".log.wf"
  LOG_FLAG = log.Ldate|log.Lmicroseconds|log.Lshortfile
)

var LOG *log.Logger
var LOGW *log.Logger
var LOGF *log.Logger

func bootstrap(resp http.ResponseWriter, r *http.Request) {

    GETs, err := url.ParseQuery(r.URL.RawQuery)
    if err != nil {
        LOGF.Println(err.Error())
    }

    var ctrlName string
    if _,ok := GETs["ctrl"]; ok {
        ctrlName = GETs["ctrl"][0]
    }

    var actName string
    if _,ok := GETs["act"]; ok {
        actName = GETs["act"][0]
    }

    err = controller.Do(ctrlName, actName, resp, r)
    if err != nil {
        LOGF.Println(err.Error())
    }
}

func main() {

    defer databases.Close()

    if fd, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744); err != nil {
        panic("log file open fail:" + LOG_FILE + " error info:" + err.Error())
    } else {
        LOG = log.New(fd, "NOTICE:", LOG_FLAG)
        defer fd.Close()
    }

    if fd, err := os.OpenFile(LOG_FILE_WF, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0744); err != nil {
        panic("log file open fail:" + LOG_FILE_WF + " error info:" + err.Error())
    } else {
        LOGW = log.New(fd, "WARNING:", LOG_FLAG)
        LOGF = log.New(fd, "FATAL:", LOG_FLAG)
        defer fd.Close()
    }

    LOG.Println("ACCESS")

    http.HandleFunc("/", bootstrap)
    http.ListenAndServe(":8181", nil)
    return
}
