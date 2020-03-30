package main

import (
	"html/template"
	"log"
	"net/http"
	"userLogin/route"
	"userLogin/route/upload"

	"github.com/gin-gonic/gin"
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
	r := gin.Default() // 返回默认的路由引擎

	// 定义静态文件地址
	// 参数1 为请求头
	// 参数2 静态文件位置
	r.Static("/assets", "./assets")

	// 自定义函数
	r.SetFuncMap(template.FuncMap{
		// 将字符串转 html
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 加载模板文件
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", route.Home)
	r.GET("/login", route.Login)
	r.POST("/login", route.PostLogin)
	r.POST("/upload", upload.Upload)
	r.POST("/uploads", upload.Uploads)

	// 路由参数 param in path
	// https://gin-gonic.com/zh-cn/docs/examples/param-in-path/
	r.GET("/user/:name", func(c *gin.Context) {
		// 获取路径参数
		name := c.Param("name")

		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "User Page",
			"link":  "<a href=/>Home</a>",
			"name":  name,
		})
	})

	// 此 handler 将匹配 /user/john/ 和 /user/john/send
	// 如果没有其他路由匹配 /user/john，它将重定向到 /user/john/
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		c.JSON(http.StatusOK, gin.H{
			"name":    name,
			"actions": action,
		})
	})

	// Query
	// /search?name=js&index=1&size=10
	r.GET("/search", func(c *gin.Context) {
		// 获取查询字段 size ,默认为 10
		size := c.DefaultQuery("size", "10")
		// c.Request.URL.Query().Get("index") 的一种快捷方式
		index := c.Query("index")
		name := c.Query("name")

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"index": index,
			"size":  size,
		})
	})

	// Form 参数
	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := c.PostForm("name")
		tel := c.PostForm("telphone")
		pwd := c.PostForm("passwd")

		// 数据格式验证
		if len(tel) != 11 {
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"code": 422, "msg": "手机号必须为11位"},
			)
			return
		}

		if len(pwd) == 0 {
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"code": 422, "msg": "密码不能为空"},
			)
			return
		}

		if len(name) == 0 {
			c.JSON(
				http.StatusUnprocessableEntity,
				gin.H{"code": 422, "msg": "用户名不能为空"},
			)
			return
		}

		// 打印结果
		log.Println(name, tel, pwd)

		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})

	r.POST("api/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "提交成功",
		})
	})

	r.GET("api/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get book",
		})
	})

	r.DELETE("api/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete book",
		})
	})

	r.Run(":3000")
}
