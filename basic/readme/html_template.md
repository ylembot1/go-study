### 1. 基本变量输出

**语法规则**：

- 使用 `{{.FieldName}}` 输出结构体字段
- 使用 `{{.}}` 输出整个对象
- **自动转义**：所有内容默认会被 HTML 转义（防止 XSS）

```go
package main

import (
    "html/template"
    "os"
)

func main() {
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
```

**输出结果**：

```html
<!DOCTYPE html>
<html>
<body>
  <h1>User Profile</h1>
  <!-- 输出字段 -->
  <p>Name: Alice<script></p>
  
  <!-- 输出未转义内容（危险！） -->
  <p>Raw: {Alice<script> <alice@example.com>}</p>
  
  <!-- 自动转义示例 -->
  <p>Email: <alice@example.com></p> 
</body>
</html>
```

### 2. 条件语句（if/else）

**语法规则**：

```handlebars
{{if .Condition}}
  <!-- 条件为真时显示 -->
{{else if .OtherCondition}}
  <!-- 其他条件 -->
{{else}}
  <!-- 条件为假时显示 -->
{{end}}
```

**示例**：

```go
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
```

**输出结果**：

```html
<p>Welcome back, Bob!</p>
<button>Admin Panel</button>
```

### 3. 循环遍历（range）

**语法规则**：

```handlebars
{{range $index, $element := .Slice}}
  <!-- 循环体 -->
{{else}}
  <!-- 当切片为空时显示 -->
{{end}}
```

**示例**：

```go
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
```

**输出结果**：

```html
<h3>Product List</h3>
<table>
  <tr>
    <td>0</td>
    <td>Laptop</td>
    <td>$1299.99</td>
  </tr>
  <tr>
    <td>1</td>
    <td>Phone</td>
    <td>$799.50</td>
  </tr>
  <tr>
    <td>2</td>
    <td>Tablet</td>
    <td>$399.00</td>
  </tr>
</table>
```

### 4. 函数调用（自定义模板函数）

Go 的 `html/template` 支持在模板中调用自定义函数，这极大增强了模板的灵活性。常见用途包括：字符串处理、格式化、条件判断等。

#### 基本用法

1. **定义函数映射**：通过 `template.FuncMap` 注册自定义函数。
2. **在模板中调用**：`{{funcName arg1 arg2}}` 或通过管道 `|` 传递参数。
3. **注意安全**：涉及 HTML 输出时，需用 `template.HTML` 等类型标记为安全，否则会被自动转义。

#### 示例与说明

```go
import (
    "html/template"
    "os"
    "strings"
)

func main() {
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
```

**输出结果**：

```html
<p>2 + 3 = 5</p>
<p>HELLO <B>WORLD</B></p>
<p><EM>LIMITED TIME</EM> OFFER!</p>
```

#### 说明与易错点

- **安全问题**：只有用 `template.HTML`/`template.JS`/`template.URL` 标记的内容才不会被转义。不要对用户输入直接用 `safeHTML`，防止 XSS。
- **函数参数**：模板函数参数类型要和模板传入类型一致，否则渲染时报错。
- **链式管道**：管道前后类型要兼容，否则会 panic。
- **注册时机**：必须在 `Parse` 之前注册函数映射。

#### 实际开发建议

- 推荐将常用的格式化、字符串处理、时间处理等函数注册到 FuncMap。
- 对于需要输出 HTML 的内容，优先保证内容安全，慎用 `safeHTML`。
- 可以将函数注册到全局模板，便于多页面复用。

---

### 5. 模板嵌套与布局系统（多文件模板、继承与复用）

在实际项目中，页面通常有统一的布局（如 header、footer、导航栏），而内容部分会变化。Go 的模板系统通过“模板嵌套”和“块”机制实现页面复用和继承。

#### 核心概念

- **define**：定义一个可复用的模板块（`{{define "name"}}...{{end}}`）。
- **template**：插入另一个模板（`{{template "name" .}}`）。
- **block**：定义可被覆盖的区域（`{{block "name" .}}默认内容{{end}}`）。
- **ParseFiles/ParseGlob**：加载多个模板文件。

#### 推荐文件结构

```
templates/
├── base.html    # 主布局（定义头部、脚部、内容块）
├── header.html  # 可选：头部块
├── footer.html  # 可选：脚部块
└── home.html    # 具体页面内容
```

#### base.html（主布局）

```html
<!DOCTYPE html>
<html>
<head>
  <title>{{block "title" .}}Default Title{{end}}</title>
</head>
<body>
  <header>{{block "header" .}}<h1>Default Header</h1>{{end}}</header>
  <main>{{block "content" .}}No content{{end}}</main>
  <footer>{{block "footer" .}}<p>&copy; 2023 My Company</p>{{end}}</footer>
</body>
</html>
```

#### home.html（页面内容）

```html
{{define "title"}}Home Page{{end}}

{{define "header"}}
  <h1>Welcome to {{.SiteName}}</h1>
{{end}}

{{define "content"}}
  <section>
    <h2>Special Offers</h2>
    <p>{{.PromoText | safeHTML}}</p>
  </section>
{{end}}
```

#### Go 代码加载与渲染

```go
import (
    "html/template"
    "os"
)

func main() {
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
```

**输出结果**：

```html
<!DOCTYPE html>
<html>
<head>
  <title>Home Page</title>
</head>
<body>
  <header><h1>Welcome to SuperStore</h1></header>
  <main>
    <section>
      <h2>Special Offers</h2>
      <p><b>50% off</b> all items this week!</p>
    </section>
  </main>
  <footer><p>&copy; 2023 My Company</p></footer>
</body>
</html>
```

#### 说明与易错点

- **模板入口**：`ExecuteTemplate` 的第一个参数应为布局模板（如 `base.html`），这样内容块才会被正确填充。
- **block机制**：`block` 定义的区域可被子模板 `define` 覆盖，实现类似继承。
- **数据传递**：`{{template "name" .}}` 的 `.` 代表当前数据上下文，注意数据结构要兼容。
- **文件加载顺序**：`ParseFiles` 加载顺序影响模板查找，建议主布局放在第一个。
- **调试技巧**：可用 `t.DefinedTemplates()` 查看已加载模板名。

#### 实际开发建议

- 推荐所有页面都通过主布局（base.html）渲染，保证统一风格。
- 公共区块（如 header/footer）可单独拆分文件，便于维护。
- 复杂项目可用 `block` + `define` 实现多级继承和复用。
- 充分利用自定义函数（如 `safeHTML`）提升模板表达力。

---

### 关键安全特性

1. **自动上下文感知转义**：

   ```go
   tmpl := `
   <!-- 根据位置自动转义 -->
   <div>{{.UserInput}}</div>          <!-- HTML 转义 -->
   <script>var str = "{{.UserInput}}"</script>  <!-- JS 转义 -->
   <a href="/search?q={{.UserInput}}">Search</a>  <!-- URL 转义 -->
   `
   ```

2. **手动安全标记**：

   ```go
   funcMap := template.FuncMap{
       "safeJS": func(s string) template.JS {
           return template.JS(s)
       },
       "safeURL": func(s string) template.URL {
           return template.URL(s)
       },
   }
   ```

### 最佳实践总结

1. **文件组织**：

   ```go
   // 加载整个目录
   t := template.Must(template.ParseGlob("templates/*.html"))
   
   // 明确指定文件
   t := template.Must(template.ParseFiles(
       "templates/base.html",
       "templates/header.html",
       "templates/home.html",
   ))
   ```

2. **错误处理**：

   ```go
   t, err := template.New("page").Parse(tmpl)
   if err != nil {
       log.Fatalf("Parse error: %v", err)
   }
   ```

3. **调试技巧**：

   ```go
   // 打印生成的模板
   var buf bytes.Buffer
   t.Execute(&buf, data)
   fmt.Println(buf.String())
   ```

这些规则和示例覆盖了实际开发中 95% 的使用场景。建议从简单模板开始，逐步增加复杂度，特别注意模板的自动转义特性，这是 `html/template` 区别于普通模板引擎的核心安全功能。