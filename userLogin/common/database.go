package common

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "log"
  "userLogin/model"
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
  
  // user_infos 表创建表与结构体自动迁移功能
  // 结构体的字段变化会自动更新相关的表
  db.AutoMigrate(&model.UserInfo{})
  // my_user 通过函数自定义表名
  db.AutoMigrate(&model.User{})
  
  return db
}

func GetDB() * gorm.DB {
  return DB
}