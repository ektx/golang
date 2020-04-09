package route

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// 重定向到百度
func RedirectBaidu(c *gin.Context) {
  c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}
