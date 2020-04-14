package common

import (
  "github.com/jinzhu/gorm"
  "log"
)

var DB * gorm.DB

func InitDB() * gorm.DB  {
  // 连接 mysql
  // https://gorm.io/zh_CN/docs/connecting_to_the_database.html
  db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
  
  if err != nil {
    log.Println(err)
    panic(err)
  } else {
    log.Println("连接数据库成功")
  }
  
  // 禁用表的复数形式
  // db.SingularTable(true)
  
  DB = db
  
  return db
}

func GetDB() * gorm.DB {
  return DB
}