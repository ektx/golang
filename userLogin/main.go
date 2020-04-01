package main

import (
	"log"
	"userLogin/route"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=host.docker.internal port=5432 user=postgres dbname=postgres password=123456 sslmode=disable")

	if err != nil {
		log.Println(err)
	} else {
		log.Println("连接数据库成功")
	}
	// 延迟关闭数据库
	defer db.Close()

	// 创建一个默认路由
	//r := gin.Default() // 返回默认的路由引擎
	r := route.Router()

	r.Run(":3000")
}
