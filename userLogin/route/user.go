package route

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "log"
  "net/http"
  "userLogin/model"
)

func addUser(c *gin.Context) {
  
  var u model.UserInfo
  ageErrInfo := map[string]string {
    "required": "年龄不能为空",
    "gte": "年龄不能小于 0",
    "lte": "年龄不能大于 130",
  }
  log.Print(1, &u)
  
  // 绑定值
  if err := c.ShouldBind(&u); err != nil {
    fmt.Printf("%#v", err)
    var msg = ""
    
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
  
      switch err.StructField() {
        case "Name": msg = "用户名不能为空"
        case "Gender": msg = "性别不能为空"
        case "Age": msg = ageErrInfo[err.Tag()]
      }
      fmt.Println(msg)
  
      c.JSON(http.StatusBadRequest, gin.H{
        "msg": msg,
      })
      break
    }
    return
  }
  
  // 保存
  if err := model.SaveUserInfo(&u); err != nil {
    c.JSON(http.StatusOK, gin.H{
      "msg": "添加失败",
    })
  } else {
    c.JSON(http.StatusOK, gin.H{
      "msg": "添加成功",
    })
  }
}