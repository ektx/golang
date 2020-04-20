package main

import (
	"fmt"
	"userLogin/common"
	"userLogin/model"
	"userLogin/route"
)



func main() {
	db := common.InitDB()
	// 延迟关闭数据库
	defer db.Close()
	// 运行表与结构体
	model.AutoMigrateModels()
	
	// 使用 User 结构体创建一个名为 my_test_tabel 表
	db.Table("my_test_table").CreateTable(&model.User{})
	
	// 测试 创建一个数据
	u1 := model.UserInfo{ID: 1, Name: "小布丁2", Age: 1, Gender: "男"}
	// save data
	db.Create(&u1)
	
	// 查询功能
	var u model.UserInfo
	db.First(&u)
	fmt.Printf("u: %#v\n", u)
	
	// 更新功能
	db.Model(&u).Update("name", "小布丁")
	
	// 删除
	//db.Delete(&u)
	
	// 查询所有数据
	var users []model.UserInfo
	db.Debug().Find(&users)//.Count(&count)
	fmt.Printf("all users: %#v\n", users)

	// 创建一个默认路由
	r := route.Router()

	r.Run(":3000")
}
