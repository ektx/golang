package route

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "log"
  "net/http"
  "userLogin/common"
  "userLogin/model"
)

func addUser(c *gin.Context) {
  db := common.GetDB()
  
  var u model.UserInfo
  log.Print(1, &u)
  
  // 绑定值
  if err := c.ShouldBind(&u); err != nil {
    fmt.Printf("%#v", err)
    
    for _, err := range err.(validator.ValidationErrors) {
      fmt.Println(err.Namespace())
      fmt.Println(err.Field())
      fmt.Println(err.StructNamespace())
      fmt.Println(err.StructField())
      fmt.Println(err.Tag())
      fmt.Println(err.ActualTag())
      fmt.Println(err.Kind())
      fmt.Println(err.Type())
      fmt.Println(err.Value())
      fmt.Println(err.Param())
      fmt.Println()
    }
  
  
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  // 保存
  if err := db.Debug().Create(&u).Error; err != nil {
    c.JSON(http.StatusOK, gin.H{
      "msg": "添加失败",
    })
  } else {
    c.JSON(http.StatusOK, gin.H{
      "msg": "添加成功",
    })
  }
}