package models

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
  //database, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/webapp")
  //database, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/webapp")

  dsn := "root:root@tcp(localhost:3306)/webapp?charset=utf8mb4&parseTime=True&loc=Local"
  database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("Failed to connect to database!")
  }

  database.AutoMigrate(&Task{})

  DB = database
}