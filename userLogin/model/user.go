package model

import (
  "database/sql"
  "github.com/jinzhu/gorm"
  "time"
  "userLogin/common"
)

// 定义一个模型
type User struct {
  gorm.Model // 内嵌 gorm.Model
  Name 					string
  Age 					sql.NullFloat64 // 零值类型
  Birthday 			*time.Time
  Email 				string 			`gorm:"type:varchar(120);unique_index"` // 长度120，唯一
  Role 					string 			`gorm:"size:255"` // 设置角色字段大小为255
  MemberNumber 	*string 		`gorm:"unique;not null"` // 设置会员号唯一且不能为空
  Num 					int 				`gorm:"AUTO_INCREMENT"` // 设置为自增类型
  Address 			string 			`gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe 			int 				`gorm:"-"` // 忽略本字段
}


// 自定义表名, 将 Model User的表名由默认 users => my_user
func (User) TableName() string {
  return "my_user"
}

// 定义一个用户结构体
type UserInfo struct {
  ID 			uint
  Name 		string  `json:"name" form:"name" gorm:"unique;not null" binding:"required"`
  Gender 	string  `json:"gender" form:"gender" gorm:"not null" binding:"required"`
  Age 		int     `json:"age" form:"age" gorm:"not null" binding:"required,gte=0,lte=130"`
}


func SaveUserInfo (user *UserInfo) (err error) {
  db := common.GetDB()
  
  // 保存
  if err := db.Debug().Create(&user).Error; err != nil {
    return err
  }
  
  return
}

// 过滤条件
type FilterList struct {
  Name string `json:"name"`
  Gender string `json:"gender"`
  Age int `json:"age"`
}

func GetUsers () (list []FilterList, err error) {
  db := common.GetDB()

  if err := db.
    Debug().
    Table("user_infos").
    Select("name, age, gender").
    Scan(&list).Error; err != nil {
    return list, err
  }
  
  return
}