package main

import (
  "fmt"
  "time"
  "net/http"
  "database/sql"
  _"github.com/go-sql-driver/mysql"
)

const (
    USERNAME = "root"
    PASSWORD = "mickey123"
    NETWORK = "tcp"
    HOST = "localhost"
    PORT = 3306
    DATABASE = "novel"
    CHARSET = "utf8mb4"
    MaxOpenConns = 3000
    MaxIdleConns = 1000
)

type completeInfo struct {
    id int64
    act_id int32
    stage int64
    total_amount int64
    user_count int64
    status int8
    create_time int32
    update_time int32
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hi")
}

type ast struct {
  i int
}

func main(){

   var a *ast
   fmt.Println(a.i)
   return
   //http.HandleFunc("/", indexHandler)
   //http.ListenAndServe(":8181", nil)

   dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s", USERNAME, PASSWORD, NETWORK, HOST, PORT, DATABASE, CHARSET)
   db, err := sql.Open("mysql", dsn)
   if err != nil {
       return
   }
   defer db.Close()
   db.SetConnMaxLifetime(10 * time.Second)
   db.SetMaxOpenConns(MaxOpenConns)
   db.SetMaxIdleConns(MaxIdleConns)

   rows, err := db.Query("SELECT `id`, `act_id`, `stage`, `total_amount`, `user_count`, `status`, `create_time`, `update_time` FROM `novel_act_complete_info` WHERE `id` = ?", 1)

   if err != nil {
       fmt.Println(err)
       return
   }
   
   for rows.Next() {

       var srow completeInfo
//       columns,_ := rows.Columns()
//       cvalue := make([][]byte,len(columns))
       err := rows.Scan(&srow)
       if err != nil {
           fmt.Println(err)
           return
       }
       
//       cinfo := make(map[string][]byte, len(columns))
//       for k, v := range columns {
//           fmt.Println(k)
//       }
       fmt.Printf("%+v", &srow)
   }
}
