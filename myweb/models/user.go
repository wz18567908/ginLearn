package models

type User struct {
   Id   int    `gorm:"primary_key"`
   Name string `gorm:"type:varchar(256);not null;"`
   Age  int
}

