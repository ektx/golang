package main

import (
	"database/sql"
	"fmt"
	"time"
	"userLogin/common"
	"userLogin/route"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义一个用户结构体
type UserInfo struct {
	ID 			uint
	Name 		string
	Gender 	string
	Age 		int
}

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

func main() {
	db := common.InitDB()
	// 延迟关闭数据库
	defer db.Close()
	
	// 创建表与结构体自动迁移功能
	// 结构体的字段变化会自动更新相关的表
	db.AutoMigrate(&UserInfo{})
	
	db.AutoMigrate(&User{})
	
	// 使用 User 结构体创建一个名为 my_test_tabel 表
	db.Table("my_test_table").CreateTable(&User{})
	
	// 测试 创建一个数据
	u1 := UserInfo{ID: 1, Name: "小布丁2", Age: 1, Gender: "男"}
	// save data
	db.Create(&u1)
	
	// 查询功能
	var u UserInfo
	db.First(&u)
	fmt.Printf("u: %#v\n", u)
	
	// 更新功能
	db.Model(&u).Update("name", "小布丁")
	
	// 删除
	//db.Delete(&u)
	
	// 查询所有数据
	var users []UserInfo
	db.Debug().Find(&users)//.Count(&count)
	fmt.Printf("all users: %#v\n", users)

	// 创建一个默认路由
	r := route.Router()

	r.Run(":3000")
}
