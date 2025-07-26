# Swagger 学习指南

## 目录
- [Swagger 简介](#swagger-简介)
- [工具使用方法](#工具使用方法)
- [Swagger 基本语法](#swagger-基本语法)
- [Swagger YAML 使用](#swagger-yaml-使用)
- [导入到 Postman](#导入到-postman)
- [项目结构](#项目结构)
- [常见问题](#常见问题)

## Swagger 简介

### 什么是 Swagger？
Swagger（现在称为 OpenAPI）是一个用于设计、构建、记录和使用 RESTful API 的强大工具集。它提供了一种标准化的方式来描述 API，使得开发者和用户能够更好地理解和使用 API。

### 主要特性
- **API 文档自动生成**：从代码注释自动生成 API 文档
- **交互式文档**：提供可交互的 API 测试界面
- **代码生成**：支持多种编程语言的客户端代码生成
- **标准化**：基于 OpenAPI 规范，确保文档的一致性
- **团队协作**：便于前后端开发者的沟通和协作

### 在 Go 项目中的优势
- 减少手动编写 API 文档的工作量
- 确保文档与代码的同步更新
- 提供统一的 API 测试界面
- 支持多种输出格式（JSON、YAML、HTML）

## 工具使用方法

### 1. 安装 Swagger 工具

```bash
# 安装 swag 命令行工具
go install github.com/swaggo/swag/cmd/swag@latest

# 验证安装
swag --version
```

### 2. 项目依赖

```bash
# 安装必要的依赖包
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

### 3. 生成 Swagger 文档

```bash
# 在项目根目录执行
swag init

# 指定输出目录
swag init -o ./docs

# 指定主文件
swag init -g main.go
```

### 4. 运行项目

```bash
# 编译项目
go build -o swagger-app main.go

# 运行服务
./swagger-app

# 或者直接运行
go run main.go
```

### 5. 访问 Swagger UI

启动服务后，访问以下地址：
- Swagger UI: http://localhost:8080/swagger/index.html
- API 文档 JSON: http://localhost:8080/swagger/doc.json
- API 文档 YAML: http://localhost:8080/swagger/doc.yaml

## Swagger 基本语法

### 1. 主文档注释

```go
// @title           Swagger Demo API
// @version         1.0
// @description     This is a sample server for Swagger documentation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
```

### 2. 函数注释语法

```go
// FunctionName godoc
// @Summary      简短描述
// @Description  详细描述
// @Tags         tags
// @Accept       json
// @Produce      json
// @Param        paramName  paramType  dataType  required  description
// @Success      200        {object}   ResponseType
// @Failure      400        {object}   ErrorResponse
// @Router       /path [method]
func FunctionName(c *gin.Context) {
    // 函数实现
}
```

### 3. 参数类型

| 参数类型 | 说明 | 示例 |
|---------|------|------|
| `query` | URL 查询参数 | `?page=1&size=10` |
| `path` | URL 路径参数 | `/users/{id}` |
| `body` | 请求体参数 | JSON 数据 |
| `header` | 请求头参数 | `Authorization: Bearer token` |
| `formData` | 表单数据 | `multipart/form-data` |

### 4. 数据类型

```go
// 基本类型
@Param page query int false "页码" default(1)
@Param name path string true "用户名"
@Param user body CreateUserRequest true "用户信息"

// 数组类型
@Param tags query []string false "标签列表"

// 对象类型
@Success 200 {object} User
@Success 200 {array} User
```

### 5. 响应定义

```go
// 成功响应
@Success 200 {object} User "成功返回用户信息"
@Success 201 {object} User "成功创建用户"

// 错误响应
@Failure 400 {object} ErrorResponse "请求参数错误"
@Failure 404 {object} ErrorResponse "用户不存在"
@Failure 500 {object} ErrorResponse "服务器内部错误"
```

### 6. 模型定义

```go
// User 用户模型
type User struct {
    ID        uint      `json:"id" example:"1"`
    Name      string    `json:"name" example:"张三"`
    Email     string    `json:"email" example:"zhangsan@example.com"`
    Age       int       `json:"age" example:"25"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

## Swagger YAML 使用

### 1. YAML 文件结构

生成的 `swagger.yaml` 文件包含完整的 API 规范：

```yaml
swagger: "2.0"
info:
  title: "Swagger Demo API"
  version: "1.0"
  description: "This is a sample server for Swagger documentation."
  contact:
    name: "API Support"
    url: "http://www.swagger.io/support"
    email: "support@swagger.io"
  license:
    name: "Apache 2.0"
    url: "http://www.apache.org/licenses/LICENSE-2.0.html"
host: "localhost:8080"
basePath: "/api/v1"
schemes:
  - "http"
paths:
  /users:
    get:
      summary: "获取用户列表"
      description: "获取所有用户的列表，支持分页"
      tags:
        - "users"
      parameters:
        - name: "page"
          in: "query"
          description: "页码"
          required: false
          type: "integer"
          default: 1
        - name: "per_page"
          in: "query"
          description: "每页数量"
          required: false
          type: "integer"
          default: 10
      responses:
        "200":
          description: "OK"
          schema:
            $ref: "#/definitions/model.UserListResponse"
        "400":
          description: "Bad Request"
          schema:
            type: "object"
            additionalProperties: true
```

### 2. YAML 的优势

- **可读性强**：比 JSON 更易读，支持注释
- **标准化**：符合 OpenAPI 规范
- **工具支持**：被大多数 API 工具支持
- **版本控制友好**：便于 Git 管理

### 3. 使用 YAML 的场景

- **API 设计**：在设计阶段定义 API 规范
- **代码生成**：生成客户端代码
- **文档共享**：与团队成员共享 API 文档
- **工具集成**：导入到其他 API 工具中

## 导入到 Postman

### 方法一：直接导入 YAML 文件

1. **打开 Postman**
2. **点击 "Import" 按钮**
3. **选择 "File" 标签**
4. **上传 `swagger.yaml` 文件**
5. **点击 "Import" 确认**

### 方法二：通过 URL 导入

1. **在 Postman 中点击 "Import"**
2. **选择 "Link" 标签**
3. **输入 Swagger 文档 URL**：
   ```
   http://localhost:8080/swagger/doc.yaml
   ```
4. **点击 "Continue" 和 "Import"**

### 方法三：导入 JSON 格式

1. **访问 Swagger JSON 端点**：
   ```
   http://localhost:8080/swagger/doc.json
   ```
2. **复制 JSON 内容**
3. **在 Postman 中粘贴并导入**

### Postman 导入后的配置

1. **环境变量设置**：
   - 创建环境变量 `base_url` = `http://localhost:8080`
   - 在请求中使用 `{{base_url}}/api/v1/users`

2. **认证配置**：
   - 如果 API 需要认证，在 Collection 设置中配置
   - 支持 Bearer Token、API Key 等认证方式

3. **测试脚本**：
   - 为每个请求添加测试脚本
   - 验证响应状态码和数据格式

## 项目结构

```
swagger/
├── main.go              # 主程序入口
├── handler/
│   └── user.go         # 用户相关处理器
├── model/
│   └── user.go         # 数据模型定义
├── docs/
│   ├── docs.go         # 生成的 Swagger 文档
│   ├── swagger.json    # JSON 格式文档
│   └── swagger.yaml    # YAML 格式文档
└── README.md           # 项目说明文档
```

### 关键文件说明

- **main.go**：程序入口，配置路由和 Swagger
- **handler/user.go**：包含 Swagger 注释的 API 处理器
- **model/user.go**：数据模型定义，用于生成 API 文档
- **docs/**：自动生成的 Swagger 文档文件

## 常见问题

### 1. "no swag has yet been registered" 错误

**问题**：访问 Swagger 文档时出现此错误

**解决方案**：
```go
// 在 main.go 中导入 docs 包
import _ "your-project/docs"
```

### 2. 文档不更新

**问题**：修改代码后 Swagger 文档没有更新

**解决方案**：
```bash
# 重新生成文档
swag init

# 重启服务
go run main.go
```

### 3. 模型定义不显示

**问题**：自定义模型在 Swagger 文档中不显示

**解决方案**：
- 确保模型有正确的 JSON 标签
- 在注释中正确引用模型
- 检查模型是否被使用

### 4. 参数验证不生效

**问题**：Swagger 文档中的参数验证不生效

**解决方案**：
- 使用 `binding` 标签进行验证
- 在处理器中手动验证参数
- 确保验证规则正确

### 5. 认证配置

**问题**：如何在 Swagger 中配置认证

**解决方案**：
```go
// 在 main.go 中添加认证配置
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
```

## 最佳实践

1. **保持文档同步**：每次修改 API 后及时更新文档
2. **使用有意义的描述**：为每个 API 提供清晰的描述
3. **合理分组**：使用 Tags 对 API 进行分组
4. **提供示例**：为请求和响应提供示例数据
5. **版本管理**：使用版本号管理 API 变更
6. **测试覆盖**：确保所有 API 都有对应的测试用例

## 参考资源

- [Swagger 官方文档](https://swagger.io/docs/)
- [OpenAPI 规范](https://swagger.io/specification/)
- [gin-swagger 文档](https://github.com/swaggo/gin-swagger)
- [swag 工具文档](https://github.com/swaggo/swag) 