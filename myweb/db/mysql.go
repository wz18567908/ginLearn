package db

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
//    "log"
    m "ginLearn/myweb/models"
    "fmt"
)

var db *gorm.DB
var user *m.User
var job *m.LSFJobReq

func InitDB() {
        fmt.Println("InitDB...")
    createTables(newConn())
}

func newConn() *gorm.DB {
        db, _ = gorm.Open("mysql", "root:111111@tcp(192.168.0.69:12345)/test?parseTime=true")
//    defer db.Close()

//        if err != nil {
//            log.Fatal(err.Error())
//        }
fmt.Println("new conn...")
    return db
}

func createTables(*gorm.DB) {
    if !db.HasTable(user) {
        if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(user).Error; err != nil {
            panic(err)
        }
    }
    if !db.HasTable(job) {
        if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(job).Error; err != nil {
            panic(err)
        }
    }
}

func InsertUser(u *m.User) {
    fmt.Println("插入用户...")
    db.Create(u)
//    defer db.Close()
}

func InsertJob(job *m.LSFJobReq) {
    fmt.Println("插入作业数据...")
    db.Create(job)
}
