# Go语言 defer、panic、recover 完全指南

## 1. defer 关键字

### 1.1 基本概念

`defer` 用于延迟函数调用，确保函数在程序执行结束时（通常是函数返回前）被调用，常用于资源清理。

### 1.2 基本语法

```go
func example() {
    defer fmt.Println("最后执行")
    fmt.Println("先执行")
    // 函数结束时，defer语句会按LIFO（后进先出）顺序执行
}
```

### 1.3 defer 执行顺序

```go
func deferOrder() {
    defer fmt.Println("第一个defer")
    defer fmt.Println("第二个defer")
    defer fmt.Println("第三个defer")
    
    fmt.Println("函数体执行")
    
    // 输出顺序：
    // 函数体执行
    // 第三个defer
    // 第二个defer
    // 第一个defer
}
```

### 1.4 常见使用场景

#### 1.4.1 文件操作

```go
func fileOperation() {
    // 打开文件
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("打开文件失败:", err)
        return
    }
    
    // 确保文件被关闭
    defer file.Close()
    
    // 读取文件内容
    data := make([]byte, 1024)
    _, err = file.Read(data)
    if err != nil {
        fmt.Println("读取文件失败:", err)
        return
    }
    
    fmt.Println("文件内容:", string(data))
    // 函数结束时，file.Close() 会自动调用
}
```

#### 1.4.2 数据库连接

```go
func databaseOperation() {
    // 连接数据库
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        fmt.Println("连接数据库失败:", err)
        return
    }
    
    // 确保关闭数据库连接
    defer db.Close()
    
    // 执行查询
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        fmt.Println("查询失败:", err)
        return
    }
    
    // 确保关闭查询结果
    defer rows.Close()
    
    // 处理查询结果
    for rows.Next() {
        // 处理每一行数据
    }
}
```

#### 1.4.3 互斥锁

```go
func mutexExample() {
    var mu sync.Mutex
    
    mu.Lock()
    defer mu.Unlock() // 确保锁被释放
    
    // 临界区代码
    fmt.Println("在临界区内执行")
}
```

#### 1.4.4 HTTP响应

```go
func httpHandler(w http.ResponseWriter, r *http.Request) {
    // 设置响应头
    w.Header().Set("Content-Type", "application/json")
    
    // 确保响应被发送
    defer func() {
        if err := recover(); err != nil {
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    }()
    
    // 处理请求
    response := map[string]string{"status": "success"}
    json.NewEncoder(w).Encode(response)
}
```

### 1.5 defer 的注意事项

#### 1.5.1 参数求值时机

```go
func deferParameter() {
    i := 0
    defer fmt.Println("defer中的i:", i) // i的值在defer时就被确定了
    
    i = 10
    fmt.Println("函数中的i:", i)
    
    // 输出：
    // 函数中的i: 10
    // defer中的i: 0
}
```

#### 1.5.2 闭包中的defer

```go
func deferClosure() {
    i := 0
    defer func() {
        fmt.Println("defer闭包中的i:", i) // 闭包会捕获最新的i值
    }()
    
    i = 10
    fmt.Println("函数中的i:", i)
    
    // 输出：
    // 函数中的i: 10
    // defer闭包中的i: 10
}
```

## 2. panic 关键字

### 2.1 基本概念

`panic` 用于触发程序崩溃，通常表示发生了不可恢复的错误。当 `panic` 发生时，程序会立即停止执行并开始执行 `defer` 函数。

### 2.2 基本语法

```go
func panicExample() {
    panic("这是一个panic")
    fmt.Println("这行代码不会执行")
}
```

### 2.3 常见使用场景

#### 2.3.1 参数验证

```go
func divide(a, b int) int {
    if b == 0 {
        panic("除数不能为零")
    }
    return a / b
}

func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("计算错误: %v", r)
        }
    }()
    
    result = divide(a, b)
    return
}
```

#### 2.3.2 初始化失败

```go
func initializeDatabase() {
    // 尝试连接数据库
    if !canConnectToDatabase() {
        panic("无法连接到数据库，程序无法继续")
    }
    
    // 初始化表结构
    if !initializeTables() {
        panic("无法初始化数据库表结构")
    }
    
    fmt.Println("数据库初始化成功")
}
```

#### 2.3.3 配置错误

```go
func loadConfig() *Config {
    config, err := readConfigFile("config.json")
    if err != nil {
        panic(fmt.Sprintf("加载配置文件失败: %v", err))
    }
    
    if config.Port == 0 {
        panic("配置文件中端口号无效")
    }
    
    return config
}
```

### 2.4 panic 的传播

```go
func level3() {
    panic("level3 panic")
}

func level2() {
    defer fmt.Println("level2 defer")
    level3()
}

func level1() {
    defer fmt.Println("level1 defer")
    level2()
}

func panicPropagation() {
    defer fmt.Println("main defer")
    
    fmt.Println("开始执行")
    level1()
    fmt.Println("这行不会执行")
    
    // 输出：
    // 开始执行
    // main defer
    // level1 defer
    // level2 defer
    // panic: level3 panic
}
```

## 3. recover 关键字

### 3.1 基本概念

`recover` 用于捕获 `panic`，只能在 `defer` 函数中使用。它可以让程序从 `panic` 中恢复，而不是崩溃退出。

### 3.2 基本语法

```go
func recoverExample() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("捕获到panic:", r)
        }
    }()
    
    panic("测试panic")
    fmt.Println("这行不会执行")
}
```

### 3.3 常见使用场景

#### 3.3.1 Web服务器错误处理

```go
func webHandler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Handler panic: %v", r)
            http.Error(w, "Internal Server Error", 500)
        }
    }()
    
    // 处理HTTP请求
    processRequest(w, r)
}

func processRequest(w http.ResponseWriter, r *http.Request) {
    // 模拟可能发生panic的操作
    if r.URL.Path == "/panic" {
        panic("模拟panic")
    }
    
    fmt.Fprintf(w, "Hello, World!")
}
```

#### 3.3.2 工作协程保护

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Worker %d panic: %v", id, r)
            // 可以选择重启worker或记录错误
        }
    }()
    
    for job := range jobs {
        // 处理任务，可能发生panic
        result := processJob(job)
        results <- result
    }
}

func processJob(job int) int {
    if job == 13 {
        panic("不吉利的数字")
    }
    return job * 2
}
```

#### 3.3.3 优雅的错误处理

```go
func safeOperation() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("操作失败: %v", r)
        }
    }()
    
    // 可能发生panic的操作
    result = riskyOperation()
    return
}

func riskyOperation() string {
    // 模拟可能失败的操作
    if time.Now().Unix()%2 == 0 {
        panic("随机失败")
    }
    return "操作成功"
}
```

## 4. 实际应用示例

### 4.1 文件处理工具

```go
type FileProcessor struct {
    filePath string
    file     *os.File
}

func NewFileProcessor(path string) *FileProcessor {
    return &FileProcessor{filePath: path}
}

func (fp *FileProcessor) Process() error {
    // 打开文件
    file, err := os.Open(fp.filePath)
    if err != nil {
        return fmt.Errorf("打开文件失败: %v", err)
    }
    fp.file = file
    
    // 确保文件被关闭
    defer func() {
        if fp.file != nil {
            fp.file.Close()
        }
    }()
    
    // 处理文件内容
    return fp.processContent()
}

func (fp *FileProcessor) processContent() error {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("处理文件内容时发生panic: %v", r)
        }
    }()
    
    // 读取文件内容
    content, err := io.ReadAll(fp.file)
    if err != nil {
        return fmt.Errorf("读取文件失败: %v", err)
    }
    
    // 处理内容（可能发生panic）
    processed := fp.transformContent(content)
    
    // 写入处理后的内容
    return fp.writeProcessedContent(processed)
}

func (fp *FileProcessor) transformContent(content []byte) []byte {
    // 模拟可能发生panic的转换操作
    if len(content) == 0 {
        panic("空文件内容")
    }
    
    // 简单的转换：转换为大写
    return bytes.ToUpper(content)
}

func (fp *FileProcessor) writeProcessedContent(content []byte) error {
    outputPath := fp.filePath + ".processed"
    return os.WriteFile(outputPath, content, 0644)
}
```

### 4.2 HTTP中间件

```go
func PanicRecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("HTTP请求panic: %v", r)
                log.Printf("请求URL: %s", r.URL.String())
                log.Printf("请求方法: %s", r.Method)
                log.Printf("用户代理: %s", r.UserAgent())
                
                // 返回500错误
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        defer func() {
            log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

### 4.3 数据库事务处理

```go
func (db *Database) Transaction(fn func(*sql.Tx) error) error {
    tx, err := db.conn.Begin()
    if err != nil {
        return fmt.Errorf("开始事务失败: %v", err)
    }
    
    defer func() {
        if r := recover(); r != nil {
            // 发生panic时回滚事务
            if rbErr := tx.Rollback(); rbErr != nil {
                log.Printf("回滚事务失败: %v", rbErr)
            }
            panic(r) // 重新抛出panic
        }
    }()
    
    // 执行事务函数
    if err := fn(tx); err != nil {
        // 函数返回错误时回滚事务
        if rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("回滚事务失败: %v, 原错误: %v", rbErr, err)
        }
        return err
    }
    
    // 提交事务
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("提交事务失败: %v", err)
    }
    
    return nil
}

// 使用示例
func transferMoney(db *Database, from, to int, amount float64) error {
    return db.Transaction(func(tx *sql.Tx) error {
        // 扣除转出账户余额
        _, err := tx.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ?", amount, from)
        if err != nil {
            return err
        }
        
        // 增加转入账户余额
        _, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, to)
        if err != nil {
            return err
        }
        
        return nil
    })
}
```

### 4.4 资源池管理

```go
type ResourcePool struct {
    resources chan *Resource
    factory   func() *Resource
}

type Resource struct {
    id   int
    data string
}

func NewResourcePool(factory func() *Resource, size int) *ResourcePool {
    pool := &ResourcePool{
        resources: make(chan *Resource, size),
        factory:   factory,
    }
    
    // 预创建资源
    for i := 0; i < size; i++ {
        pool.resources <- factory()
    }
    
    return pool
}

func (p *ResourcePool) Get() *Resource {
    select {
    case resource := <-p.resources:
        return resource
    default:
        // 池中没有可用资源，创建新的
        return p.factory()
    }
}

func (p *ResourcePool) Put(resource *Resource) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("归还资源到池中时发生panic: %v", r)
            // 可以选择销毁资源或记录错误
        }
    }()
    
    select {
    case p.resources <- resource:
        // 成功归还到池中
    default:
        // 池已满，销毁资源
        p.destroyResource(resource)
    }
}

func (p *ResourcePool) destroyResource(resource *Resource) {
    // 清理资源的逻辑
    log.Printf("销毁资源: %d", resource.id)
}
```

## 5. 最佳实践

### 5.1 defer 最佳实践

1. **尽早使用defer**：在资源获取后立即使用defer
2. **避免在循环中使用defer**：可能导致资源泄漏
3. **使用命名返回值**：defer可以修改命名返回值

```go
func bestPracticeDefer() (result string, err error) {
    defer func() {
        if err != nil {
            result = "" // 清空结果
        }
    }()
    
    // 处理逻辑
    return "success", nil
}
```

### 5.2 panic 最佳实践

1. **只在真正不可恢复的情况下使用panic**
2. **在包边界处理panic**：不要让panic传播到包外部
3. **提供有意义的panic消息**

```go
func validateConfig(config *Config) {
    if config == nil {
        panic("配置对象不能为nil")
    }
    
    if config.Port <= 0 || config.Port > 65535 {
        panic(fmt.Sprintf("无效的端口号: %d", config.Port))
    }
}
```

### 5.3 recover 最佳实践

1. **只在适当的地方使用recover**：如HTTP处理器、工作协程
2. **记录panic信息**：帮助调试
3. **不要隐藏panic**：除非确实需要恢复

```go
func safeHandler(handler func()) {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Handler panic: %v", r)
            log.Printf("Stack trace: %s", debug.Stack())
            // 可以选择重启服务或发送告警
        }
    }()
    
    handler()
}
```

## 6. 常见陷阱和注意事项

### 6.1 defer 陷阱

```go
// 陷阱1：defer在循环中
func deferInLoop() {
    for i := 0; i < 3; i++ {
        defer fmt.Println(i) // 输出：2, 1, 0
    }
}

// 陷阱2：defer函数参数求值
func deferParameterTrap() {
    i := 0
    defer fmt.Println(i) // 输出0，不是10
    i = 10
}

// 陷阱3：defer在return之后
func deferAfterReturn() {
    defer fmt.Println("defer")
    return
    defer fmt.Println("这个defer不会执行") // 编译错误
}
```

### 6.2 panic/recover 陷阱

```go
// 陷阱1：recover不在defer中
func recoverTrap() {
    if r := recover(); r != nil { // 这不会工作
        fmt.Println("recovered:", r)
    }
    panic("test")
}

// 陷阱2：recover在错误的goroutine中
func recoverWrongGoroutine() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                fmt.Println("recovered in goroutine:", r)
            }
        }()
        panic("goroutine panic")
    }()
    
    // 主goroutine的panic不会被上面的recover捕获
    time.Sleep(time.Second)
    panic("main panic")
}
```

## 7. 总结

### 7.1 使用建议

1. **defer**：用于资源清理，确保资源正确释放
2. **panic**：用于不可恢复的错误，避免程序继续执行
3. **recover**：用于从panic中恢复，保持程序稳定运行

### 7.2 组合使用模式

```go
func robustFunction() (result string, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("函数panic: %v", r)
        }
    }()
    
    // 可能发生panic的操作
    result = riskyOperation()
    return
}
```

这三个关键字是Go语言错误处理的重要组成部分，合理使用可以让程序更加健壮和可靠！ 