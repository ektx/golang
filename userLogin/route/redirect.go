package route

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// 重定向到百度
func RedirectBaidu(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

func RedirectEventA(c *gin.Context) {
  //跳转到事件B
  c.Request.URL.Path = "/RedirectEventB"
}

func RedirectEventB(c *gin.Context)  {
  c.JSON(http.StatusOK, gin.H{
    "msg": "来至事件 B",
  })
}