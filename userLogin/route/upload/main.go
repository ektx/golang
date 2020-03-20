package upload

import (
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// Upload 单个文件上传
func Upload(c *gin.Context) {
	file, err := c.FormFile("file")

	// 上传失败
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 20000,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	// 设置保存路径
	savePath := path.Join("./assets", file.Filename)

	log.Println(savePath)

	// 监听上传保存错误信息
	err = c.SaveUploadedFile(file, savePath)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 20001,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	log.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "上传成功",
		"data": "",
	})
}

// Uploads 上传多个文件
func Uploads(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 20001,
			"msg":  err.Error(),
			"data": "",
		})
		return
	}

	files := form.File["file"]
	// 定义一个切片 用于返回用户上传的文件列表,默认为[]
	data := []string{}

	for index, file := range files {
		log.Println(index, file.Filename)
		// append(切片, 数据1, [...数据n])
		data = append(data, file.Filename)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "上传成功!",
		"data": data,
	})
}
