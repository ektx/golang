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
    fmt.Printf("%#v\n", err)
    var msg = ""
    
    for _, err := range err.(validator.ValidationErrors) {
      // 验证结构体字段
      // eg：Namespaceerr: UserInfo.Age
      //fmt.Println("Namespaceerr:", err.Namespace())
      //fmt.Println("Field:", err.Field())            // 结构体验证字段 eg: Age
      //fmt.Println("StructNamespace:", err.StructNamespace())  // 结构体字段，同Namespace
      //fmt.Println("StructField:", err.StructField())      // 结构体错误 tag
      //fmt.Println("Tag:", err.Tag())              // 同上
      //fmt.Println("ActualTag:", err.ActualTag())        // 结构体类型,eg: string
      //fmt.Println("Kind:", err.Kind())             // 同上
      //fmt.Println("Type:", err.Type())             // 错误值
      //fmt.Println("Value:",err.Value())            // 正确值
      //fmt.Println("Param:", err.Param())
  
      switch err.Field() {
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

func GetUsers(c * gin.Context) {
  list, err := model.GetUsers()
  
  if err != nil {
    c.JSON(http.StatusOK, gin.H{
      "msg": "查询失败",
    })
  } else {
    c.JSON(http.StatusOK, gin.H{
      "msg": "查询成功",
      "data": list,
    })
  }
}