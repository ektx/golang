package route

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"userLogin/route/upload"
	
	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "hello Gin!",
		"link":  "<a href=/user/ektx>User</a>",
	})
}

// 自定义中间件
func hasLogin (status bool) gin.HandlerFunc {
	// 连接数据库等相关操作
	
	// 返回一个闭包
	return func(c * gin.Context) {
		log.Print("来到了中间件")
		if (status) {
			c.Next()
		} else {
			c.Abort()
		}
	}
}

// 中间件设置值 demo
func computedTime (c * gin.Context) {
	log.Println("进入时间中间件")
	// 开始时间
	start := time.Now()
	
	// 添加值
	c.Set("name", "小布丁")
	
	c.Next()
	end := time.Since(start)
	
	fmt.Printf("耗时：%v\n ", end)
	log.Println("结束时间中间件")
}

// 中间件取值 Demo
func apiMid (c * gin.Context) {
	log.Println("进入API 中间件1")
	// 取值
	name, has := c.Get("name")
	
	if !has {
		name = "宝宝"
	}
	
	c.Next()
	
	log.Println("中间件传值为：", name)
	log.Println("结束 API 中间件1")
}

func apiMid2 (c * gin.Context) {
	log.Println("进入API 中间件2")
	c.Next()
	log.Println("结束 API 中间件2")
}

func Router() *gin.Engine  {
	// 默认使用了 Logger() 与 Recovery() 中中间件
	// Logger 用于将日志写入 gin.DefaultWriter,即使配制了 GIN_MODE=release
	// Recovery 用于防止服务器崩溃，当有错误时，会使用500响应
	r := gin.Default()
	
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
	
	// 全局注册中间件
	r.Use(computedTime)
	
	r.GET("/", Home) // 主页
	r.GET("/baidu", RedirectBaidu) // 重定向到百度
	r.GET("/login", Login) // 登录页面
	
	// 使用单个中间件 通过
	r.GET("/hasLogin", hasLogin(true), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})
	
	// 使用单个中间件 不通过
	r.GET("/notLogin", hasLogin(false), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "OK",
		})
	})
	r.GET("/bindUser", ShouldBind)
	r.POST("/login", PostLogin) // 登录请求
	r.POST("/upload", upload.Upload) // 单文件上传
	r.POST("/uploads", upload.Uploads) // 多文件上传
	r.POST("/bindUser", ShouldBind)
	
	r.POST("/api/addUser", addUser)			// 添加用户
	r.GET("/api/getUsers", GetUsers)   	// 获取用户
	
	// 重定向事件A
	r.GET("/redirectEventA", func (c *gin.Context) {
		//跳转到事件B
		c.Request.URL.Path = "/redirectEventB"
		r.HandleContext(c)
	})
	
	// 重定向事件B
	r.GET("/redirectEventB", func (c *gin.Context)  {
		c.JSON(http.StatusOK, gin.H{
			"msg": "来至事件 B",
		})
	})
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
	
	// 组添加中间件方式一
	apiGroup := r.Group("/api", apiMid)
	// 组添加中间件方式二
	apiGroup.Use(apiMid2)
	{
		apiGroup.POST("/book", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "提交成功",
			})
		})
		
		apiGroup.GET("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Get book",
			})
		})
		
		apiGroup.DELETE("/book", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Delete book",
			})
		})
	}
	
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404/index.tmpl", gin.H{})
	})
	
	return r
}