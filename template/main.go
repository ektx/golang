package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", indexView)
	http.HandleFunc("/user", userView)
	http.HandleFunc("/fun", funView)

	// 监听 8080 端口
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("HTTP Server start failed, error: %v", err)
	}
}

func indexView(w http.ResponseWriter, r *http.Request) {
	// 解析模板内容
	t, err := template.ParseFiles("./index.tmpl")

	// 如果解析模板出错
	if err != nil {
		// 输出错误并结束
		fmt.Println("Parse template failed, error: %v", err)
		return
	}

	// 渲染数据
	t.Execute(w, "你好")
}

// User 定义一个 user 结构体
// 注意不能用小写开头，无法读取
type User struct {
	Name   string
	Age    int
	Gender string
}

func userView(w http.ResponseWriter, r *http.Request) {
	// 解析嵌套模板
	// 先写主模板，再写引用模板
	t, err := template.ParseFiles("./user.tpl", "child.tpl")

	// 如果解析模板出错
	if err != nil {
		// 输出错误并结束
		fmt.Println("Parse template failed, error: %v", err)
		return
	}

	// 设置结构体数据
	u1 := User{
		Name:   "小布丁",
		Age:    1,
		Gender: "男",
	}
	// 设置一个map数据
	m1 := map[string]interface{}{
		"name":   "宝宝",
		"age":    18,
		"gender": "女",
	}
	// 定义一个数组
	hobbyList := []string{
		"抽烟",
		"喝酒",
		"烫头",
	}

	// 渲染数据
	// t.Execute(w, u1)
	t.Execute(w, map[string]interface{}{
		"u1":       u1,
		"m1":       m1,
		"hobby":    hobbyList,
		"emptyArr": []string{},
	})
}

func funView(w http.ResponseWriter, r *http.Request) {
	// 1.定义一个方法
	// name: string 名称
	// 返回2个参数 string>处理后的结果，error>错误信息，参数2必须为错误类型
	hi := func(name string) (string, error) {
		return "你好! " + name, nil
	}

	// 2.定义模板
	// 创建一个 fun.tpl 的模板对象
	t := template.New("fun.tpl")
	// 自定义解析关键字
	t.Delims("{[", "]}")
	// 在模板对象上创建一个自定义函数 hello，方法为 hi
	t.Funcs(template.FuncMap{
		"hello": hi,
	})

	// 3.解析模板
	_, err := t.ParseFiles("./fun.tpl")

	if err != nil {
		fmt.Println("解析模板出错:%v", err)
		return
	}

	t.Execute(w, "小布丁")
}
