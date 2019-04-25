package services

import (
    user "ginLearn/myweb/models"
    db "ginLearn/myweb/db"
)

func InsertUser(id int, name string, age int) {
    u := user.User{Id : id, Name : name, Age : age}
    db.InsertUser(&u)
}
