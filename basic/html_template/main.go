package main

import (
	"html/template"
	"os"
	"strings"
)

func basic_variable() {
	type User struct {
		Name  string
		Email string
	}

	tmpl := `
<!DOCTYPE html>
<html>
<body>
  <h1>User Profile</h1>
  <!-- 输出字段 -->
  <p>Name: {{.Name}}</p>
  
  <!-- 输出未转义内容（危险！） -->
  <p>Raw: {{.}}</p>
  
  <!-- 自动转义示例 -->
  <p>Email: {{.Email}}</p> 
</body>
</html>`

	// 注意：Email 包含特殊字符
	user := User{
		Name:  "Alice<script>",
		Email: "<alice@example.com>",
	}

	t := template.Must(template.New("profile").Parse(tmpl))
	t.Execute(os.Stdout, user)
}

func if_else() {
	tmpl := `
{{if .LoggedIn}}
  <p>Welcome back, {{.User.Name}}!</p>
  {{if eq .User.Role "admin"}}
    <button>Admin Panel</button>
  {{else}}
    <button>User Settings</button>
  {{end}}
{{else}}
  <a href="/login">Please login</a>
{{end}}`

	data := struct {
		LoggedIn bool
		User     struct {
			Name string
			Role string
		}
	}{
		LoggedIn: true,
		User: struct {
			Name string
			Role string
		}{Name: "Bob", Role: "admin"},
	}

	t := template.Must(template.New("if_else").Parse(tmpl))
	t.Execute(os.Stdout, data)
}

func range_loop() {
	tmpl := `
<h3>Product List</h3>
<table>
{{range $i, $p := .Products}}
  <tr>
    <td>{{$i}}</td>         <!-- 索引 -->
    <td>{{$p.Name}}</td>    <!-- 产品名 -->
    <td>\${{printf "%.2f" $p.Price}}</td> <!-- 格式化价格 -->
  </tr>
{{else}}
  <tr><td colspan="3">No products available</td></tr>
{{end}}
</table>`

	data := struct {
		Products []struct {
			Name  string
			Price float64
		}
	}{
		Products: []struct {
			Name  string
			Price float64
		}{
			{"Laptop", 1299.99},
			{"Phone", 799.50},
			{"Tablet", 399.00},
		},
	}

	t := template.Must(template.New("for_loop").Parse(tmpl))
	t.Execute(os.Stdout, data)
}

func function_call() {
	funcMap := template.FuncMap{
		// 标记为安全 HTML（慎用！仅用于信任内容）
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		// 字符串转大写
		"uppercase": strings.ToUpper,
		// 加法
		"add": func(a, b int) int { return a + b },
	}

	tmpl := `
<!-- 直接调用自定义函数 -->
<p>2 + 3 = {{add 2 3}}</p>

<!-- 管道操作：先转大写再标记为安全HTML -->
<p>{{"hello <b>world</b>" | uppercase | safeHTML}}</p>

<!-- 结构体字段管道链式处理 -->
<p>{{.Description | uppercase | safeHTML}}</p>
`

	data := struct {
		Description string
	}{
		Description: "<em>limited time</em> offer!",
	}

	t := template.Must(template.New("funcs").Funcs(funcMap).Parse(tmpl))
	t.Execute(os.Stdout, data)
}

func template_inheritance() {
	funcMap := template.FuncMap{
		"safeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
	}

	t := template.Must(template.New("").Funcs(funcMap).ParseFiles(
		"templates/base.html",
		"templates/home.html",
	))

	data := struct {
		SiteName  string
		PromoText string
	}{
		SiteName:  "SuperStore",
		PromoText: "<b>50% off</b> all items this week!",
	}

	t.ExecuteTemplate(os.Stdout, "base.html", data)
}

func main() {
	// basic_variable()

	// if_else()

	// range_loop()

	// function_call()

	template_inheritance()
}
