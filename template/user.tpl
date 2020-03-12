<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Go Template</title>
</head>
<body>
  <p>{{ . }}</p>
  {{/* 移除左空格 */}}
  <h1>{{- .u1.Name }}</h1>
  <ul>
    <li>年龄: {{ .u1.Age }}</li>
    <li>性别: {{ .u1.Gender }}</li>
  </ul>
  <h1>{{ .m1.name }}</h1>
  <ul>
    <li>年龄: {{ .m1.age }}</li>
    <li>性别: {{ .m1.gender }}</li>
  </ul>

  <h2>With</h2>
  {{with .m1}}
  <ul>
    <li>年龄: {{ .age }}</li>
    <li>性别: {{ .gender }}</li>
  </ul>
  {{end}}

  {{/* 定义变量 */}}
  {{ $number := 1000}}
  {{ $name := .m1.name}}

  <p>$number: {{$number}}</p>
  <p>$name: {{$name}}</p>

  <h2>if</h2>
  {{if $number}}
  {{$number}}
  {{else}}
  Empty
  {{end}}

  {{/* lt 小于 */}}
  {{if lt $number 9999}}
  <p>{{$number}}小于9999</p>
  {{else}}
  <p>大于9999</p>
  {{end}}

  <h2>Range</h2>
  {{range $index, $data := .hobby}}
  <p>{{$index}} - {{$data}}</p>
  {{end}}

  {{range $i,$item := .emptyArr}}
  <p>{{$i}} - {{$item}}</p>
  {{else}}
  <p>没有内容</p>
  {{end}}

  <h2>嵌套模板</h2>
  <h3>外部引入</h3>
  {{template "child.tpl"}}
  <h3>内部定义模板</h3>
  {{template "insert.tpl"}}

  {{/* 使用 define 定义内部模板 */}}
  {{define "insert.tpl"}}
  <h1>我是 define 定义模板</h1>
  {{end}}
</body>
</html>