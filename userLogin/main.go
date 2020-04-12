package main

import (
	"fmt"
	"log"
	"userLogin/route"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 定义一个用户结构体
type UserInfo struct {
	ID uint
	Name string
	Gender string
	Age int
}

func main() {
	// 连接 mysql
	// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
	db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8mb4&parseTime=True&loc=Local")
	
	if err != nil {
		log.Println(err)
		panic(err)
	} else {
		log.Println("连接数据库成功")
	}
	// 延迟关闭数据库
	defer db.Close()
	
	// 创建表与结构体自动迁移功能
	// 结构体的字段变化会自动更新相关的表
	db.AutoMigrate(&UserInfo{})
	
	// 测试 创建一个数据
	//u1 := UserInfo{ID: 1, Name: "小布丁2", Age: 1, Gender: "男"}
	// save data
	//db.Create(&u1)
	
	// 查询功能
	var u UserInfo
	db.First(&u)
	fmt.Printf("u: %#v\n", u)
	
	// 更新功能
	db.Model(&u).Update("name", "小布丁")
	
	// 删除
	db.Delete(&u)

	// 创建一个默认路由
	r := route.Router()

	r.Run(":3000")
}
