package model

import "userLogin/common"

func AutoMigrateModels () {
  db := common.GetDB()
  
  // user_infos 表创建表与结构体自动迁移功能
  // 结构体的字段变化会自动更新相关的表
  db.AutoMigrate(&UserInfo{})
  // my_user 通过函数自定义表名
  db.AutoMigrate(&User{})
}
