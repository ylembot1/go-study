在 Go 语言中，虽然没有内置的 `enum` 关键字，但可以通过以下几种惯用方式实现枚举功能：

### 1. 基础实现（使用 iota）
这是最常用的方式，通过 `const` 和 `iota` 自动生成递增的值：

```go
type Status int

const (
    Pending Status = iota  // 0
    Active                 // 1
    Deleted                // 2
    Rejected               // 3
)

// 使用方法
func main() {
    s := Active
    fmt.Println(s) // 输出: 1
}
```

### 2. 带字符串描述（推荐）
添加 `String()` 方法实现可读性更好的输出：

```go
type Status int

const (
    Pending Status = iota
    Active
    Deleted
    Rejected
)

func (s Status) String() string {
    return [...]string{"Pending", "Active", "Deleted", "Rejected"}[s]
}

func main() {
    s := Active
    fmt.Println(s) // 输出: Active
}
```

### 3. 自定义值枚举
当需要特定值时（如兼容数据库值）：

```go
type LogLevel int

const (
    Debug LogLevel = 100
    Info  LogLevel = 200
    Warn  LogLevel = 300
    Error LogLevel = 400
)
```

### 4. 安全枚举（带验证）
添加验证函数防止非法值：

```go
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) IsValid() bool {
    return d >= North && d <= West
}

func main() {
    d := Direction(5)
    if !d.IsValid() {
        fmt.Println("Invalid direction")
    }
}
```

### 5. 高级模式（带方法和属性）
为枚举类型添加额外功能：

```go
type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) Hex() string {
    return [...]string{"#FF0000", "#00FF00", "#0000FF"}[c]
}

func (c Color) Name() string {
    return [...]string{"Red", "Green", "Blue"}[c]
}

func main() {
    c := Green
    fmt.Println(c.Name()) // 输出: Green
    fmt.Println(c.Hex())  // 输出: #00FF00
}
```

### 最佳实践建议：
1. **总是定义类型**：使用 `type Status int` 而非直接使用 `int`
2. **实现 Stringer 接口**：为枚举添加 `String() string` 方法
3. **使用 iota**：大部分场景下自动递增更安全
4. **零值处理**：将第一个值作为默认/未知状态（`Unknown = iota`）
5. **分组常量**：相关枚举放在同一个 `const` 块中
6. **添加验证**：重要场景添加 `IsValid()` 方法

### 完整示例（推荐方式）：
```go
type UserRole int

const (
    RoleUnknown UserRole = iota
    RoleGuest
    RoleUser
    RoleAdmin
)

func (r UserRole) String() string {
    switch r {
    case RoleGuest:
        return "Guest"
    case RoleUser:
        return "User"
    case RoleAdmin:
        return "Admin"
    default:
        return "Unknown"
    }
}

func (r UserRole) IsValid() bool {
    return r >= RoleGuest && r <= RoleAdmin
}

// 使用示例
func main() {
    role := RoleAdmin
    fmt.Println(role)       // 输出: Admin
    fmt.Println(role == 2)  // 输出: true
}
```

这些实现方式在 Go 标准库中被广泛使用（如 `time.Weekday`、`os.FileMode` 等）。选择哪种方式取决于具体需求，但带 `String()` 方法的基础 `iota` 实现能满足大多数场景。