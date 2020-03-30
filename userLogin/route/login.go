package route

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Login(c *gin.Context) {
  c.HTML(http.StatusOK, "login/index.html", nil)
}

func PostLogin(c *gin.Context)  {
  // 获取方法一
  //username := c.PostForm("username")
  //password := c.PostForm("password")
  
  // 获取方法二
  // 取不到值时（空也是值）
  username := c.DefaultPostForm("username", "admin")
  //password := c.DefaultPostForm("password", "123456")
  
  // 方法三
  password, ok := c.GetPostForm("password")
  // 如果取不到值，则设置为123456
  // 注意：空也是值
  if !ok {
    password = "123456"
  }
  
  c.JSON(http.StatusOK, gin.H{
    "username": username,
    "password": password,
  })
}