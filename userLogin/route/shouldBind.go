package route

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
)

// tag 结构体
type UserInfo struct {
  Username string `form:"username" json:"user"`
  Password string `form:"password" json:"pwd"`
}

func ShouldBind (c *gin.Context) {
  // 声明一个UserInfo类型的变量
  var usr UserInfo
  
  // &usr 使用指针方式获取 usr 的值
  err := c.ShouldBind(&usr)
  
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "msg": err.Error(),
      "code": 400,
      "data": "",
    })
  } else {
    fmt.Printf("%#v\n", usr)
    c.JSON(http.StatusOK, gin.H{
      "code": 0,
      "msg": "成功",
      "data": "",
    })
  }
}