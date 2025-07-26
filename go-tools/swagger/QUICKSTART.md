# Swagger 快速开始指南

## 5分钟快速上手

### 1. 安装工具

```bash
# 安装 swag 命令行工具
go install github.com/swaggo/swag/cmd/swag@latest

# 验证安装
swag --version
```

### 2. 创建项目结构

```
my-api/
├── main.go
├── handler/
│   └── user.go
├── model/
│   └── user.go
└── docs/          # 自动生成
```

### 3. 编写主程序

```go
// main.go
package main

import (
    "github.com/gin-gonic/gin"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
    _ "my-api/docs" // 导入生成的文档
)

// @title           My API
// @version         1.0
// @description     This is a sample API
// @host            localhost:8080
// @BasePath        /api/v1

func main() {
    r := gin.Default()
    
    // 注册 Swagger 路由
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    
    // 你的 API 路由
    r.Run(":8080")
}
```

### 4. 编写处理器

```go
// handler/user.go
package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// GetUser godoc
// @Summary      获取用户信息
// @Description  根据用户ID获取用户详细信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "用户ID"
// @Success      200  {object}  User
// @Failure      404  {object}  ErrorResponse
// @Router       /users/{id} [get]
func GetUser(c *gin.Context) {
    // 你的业务逻辑
    c.JSON(http.StatusOK, gin.H{"message": "success"})
}
```

### 5. 定义模型

```go
// model/user.go
package model

type User struct {
    ID    int    `json:"id" example:"1"`
    Name  string `json:"name" example:"张三"`
    Email string `json:"email" example:"zhang@example.com"`
}

type ErrorResponse struct {
    Error   string `json:"error" example:"错误信息"`
    Message string `json:"message" example:"详细描述"`
}
```

### 6. 生成文档

```bash
# 在项目根目录执行
swag init

# 查看生成的文件
ls docs/
# docs.go  swagger.json  swagger.yaml
```

### 7. 运行项目

```bash
# 运行服务
go run main.go

# 访问 Swagger UI
# http://localhost:8080/swagger/index.html
```

## 常用命令

```bash
# 生成文档
swag init

# 指定输出目录
swag init -o ./docs

# 指定主文件
swag init -g main.go

# 指定解析目录
swag init -d ./

# 生成时排除某些目录
swag init --exclude ./vendor
```

## 常见注释语法

### 主文档注释
```go
// @title           API 标题
// @version         1.0
// @description     API 描述
// @host            localhost:8080
// @BasePath        /api/v1
// @contact.name   联系人
// @contact.email   邮箱
// @license.name    许可证
```

### 函数注释
```go
// FunctionName godoc
// @Summary      简短描述
// @Description  详细描述
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        name  type  required  description
// @Success      200   {object}  ResponseType
// @Failure      400   {object}  ErrorResponse
// @Router       /path [method]
```

### 参数类型
- `query`: URL 查询参数 `?page=1`
- `path`: URL 路径参数 `/users/{id}`
- `body`: 请求体参数
- `header`: 请求头参数
- `formData`: 表单数据

## 快速测试

### 1. 启动服务
```bash
go run main.go
```

### 2. 访问 Swagger UI
打开浏览器访问：http://localhost:8080/swagger/index.html

### 3. 测试 API
在 Swagger UI 中可以直接测试你的 API

### 4. 导出到 Postman
1. 访问：http://localhost:8080/swagger/doc.yaml
2. 复制 YAML 内容
3. 在 Postman 中导入

## 故障排除

### 问题1: "no swag has yet been registered"
**解决**: 确保在 main.go 中导入了 docs 包
```go
import _ "your-project/docs"
```

### 问题2: 文档不更新
**解决**: 重新生成文档并重启服务
```bash
swag init
go run main.go
```

### 问题3: 模型不显示
**解决**: 确保模型有正确的 JSON 标签
```go
type User struct {
    ID int `json:"id" example:"1"`
}
```

## 下一步

1. 查看完整的 [README.md](README.md) 文档
2. 学习更多 [示例代码](examples/swagger_examples.go)
3. 探索高级功能：认证、文件上传、复杂参数等
4. 集成到你的实际项目中

## 有用的链接

- [Swagger 官方文档](https://swagger.io/docs/)
- [gin-swagger 文档](https://github.com/swaggo/gin-swagger)
- [swag 工具文档](https://github.com/swaggo/swag) 