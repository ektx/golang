package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"title": "hello Gin!",
		"link":  "<a href=/user/ektx>User</a>",
	})
}
