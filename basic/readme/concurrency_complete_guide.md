# Go并发编程完整指南

## 目录
1. [概述](#概述)
2. [Goroutine](#goroutine)
3. [Channel](#channel)
4. [WaitGroup](#waitgroup)
5. [Mutex](#mutex)
6. [Atomic](#atomic)
7. [Timer](#timer)
8. [Ticker](#ticker)
9. [工具配合使用](#工具配合使用)
10. [最佳实践](#最佳实践)

## 概述

Go语言的并发编程模型基于CSP（Communicating Sequential Processes）理论，核心思想是"通过通信来共享内存，而不是通过共享内存来通信"。Go提供了丰富的并发原语，使得并发编程变得简单而高效。

### 核心概念
- **Goroutine**: 轻量级线程，由Go运行时管理
- **Channel**: 类型安全的通信机制
- **同步原语**: WaitGroup、Mutex、Atomic等
- **定时器**: Timer、Ticker用于时间相关操作

## Goroutine

### 基本概念
Goroutine是Go语言的轻量级线程，由Go运行时调度器管理。相比传统线程，Goroutine具有以下特点：
- 启动成本低（仅2KB栈空间）
- 可以轻松创建成千上万个
- 由Go运行时自动调度

### 基本用法

```go
// 启动一个goroutine
go func() {
    fmt.Println("Hello from goroutine")
}()

// 带参数的goroutine
go func(name string) {
    fmt.Printf("Hello %s from goroutine\n", name)
}("World")
```

### 应用场景
- **并行计算**: 将大任务分解为多个小任务并行执行
- **异步处理**: 处理I/O操作、网络请求等
- **事件驱动**: 处理用户输入、系统事件等
- **后台任务**: 定期清理、监控等

### 实际Demo

```go
func example_parallel_processing() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    results := make(chan int, len(numbers))
    
    // 并行处理每个数字
    for _, num := range numbers {
        go func(n int) {
            // 模拟耗时计算
            time.Sleep(time.Millisecond * 100)
            results <- n * n
        }(num)
    }
    
    // 收集结果
    for range numbers {
        fmt.Println(<-results)
    }
}
```

## Channel

### 基本概念
Channel是Go语言中用于goroutine间通信的管道，具有以下特性：
- 类型安全
- 阻塞式通信
- 支持缓冲
- 可以关闭

### 基本用法

#### 1. 无缓冲Channel
```go
// 创建无缓冲channel
ch := make(chan string)

// 发送数据
ch <- "hello"

// 接收数据
msg := <-ch
```

#### 2. 有缓冲Channel
```go
// 创建有缓冲channel
ch := make(chan string, 2)

// 可以发送多个数据而不阻塞
ch <- "first"
ch <- "second"

// 接收数据
fmt.Println(<-ch) // "first"
fmt.Println(<-ch) // "second"
```

#### 3. Channel方向
```go
// 只发送
func sendOnly(ch chan<- string, msg string) {
    ch <- msg
}

// 只接收
func receiveOnly(ch <-chan string) string {
    return <-ch
}
```

### 高级用法

#### 1. Select语句
```go
func example_select() {
    c1 := make(chan string)
    c2 := make(chan string)
    
    go func() {
        time.Sleep(time.Second)
        c1 <- "one"
    }()
    
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg := <-c1:
            fmt.Println("received:", msg)
        case msg := <-c2:
            fmt.Println("received:", msg)
        }
    }
}
```

#### 2. 超时处理
```go
func example_timeout() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(time.Second * 2)
        ch <- "result"
    }()
    
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-time.After(time.Second):
        fmt.Println("timeout")
    }
}
```

#### 3. 非阻塞操作
```go
func example_non_blocking() {
    ch := make(chan string)
    
    // 非阻塞发送
    select {
    case ch <- "message":
        fmt.Println("sent")
    default:
        fmt.Println("not sent")
    }
    
    // 非阻塞接收
    select {
    case msg := <-ch:
        fmt.Println("received:", msg)
    default:
        fmt.Println("no message")
    }
}
```

### 应用场景
- **数据传递**: 在goroutine间传递数据
- **同步**: 等待goroutine完成
- **事件通知**: 通知其他goroutine事件发生
- **工作池**: 分发任务给多个worker

### 实际Demo

#### 工作池模式
```go
func example_worker_pool() {
    const numJobs = 5
    const numWorkers = 3
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    
    // 启动workers
    for w := 1; w <= numWorkers; w++ {
        go func(workerId int) {
            for job := range jobs {
                fmt.Printf("worker %d processing job %d\n", workerId, job)
                time.Sleep(time.Second)
                results <- job * 2
            }
        }(w)
    }
    
    // 发送任务
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 收集结果
    for a := 1; a <= numJobs; a++ {
        fmt.Printf("result: %d\n", <-results)
    }
}
```

## WaitGroup

### 基本概念
WaitGroup用于等待一组goroutine完成，提供了简单的同步机制。

### 基本用法

```go
func example_waitgroup() {
    var wg sync.WaitGroup
    
    worker := func(id int) {
        defer wg.Done()
        fmt.Printf("Worker %d starting\n", id)
        time.Sleep(time.Second)
        fmt.Printf("Worker %d done\n", id)
    }
    
    // 启动多个goroutine
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i)
    }
    
    // 等待所有goroutine完成
    wg.Wait()
    fmt.Println("All workers done")
}
```

### 应用场景
- **批量任务**: 等待多个任务完成
- **资源清理**: 等待所有goroutine清理完成
- **数据收集**: 等待所有数据收集完成

### 实际Demo

```go
func example_parallel_download() {
    urls := []string{
        "https://example.com/1",
        "https://example.com/2",
        "https://example.com/3",
    }
    
    var wg sync.WaitGroup
    results := make([]string, len(urls))
    
    for i, url := range urls {
        wg.Add(1)
        go func(index int, url string) {
            defer wg.Done()
            // 模拟下载
            time.Sleep(time.Millisecond * 500)
            results[index] = fmt.Sprintf("Downloaded: %s", url)
        }(i, url)
    }
    
    wg.Wait()
    
    for _, result := range results {
        fmt.Println(result)
    }
}
```

## Mutex

### 基本概念
Mutex（互斥锁）用于保护共享资源，确保同一时间只有一个goroutine可以访问。

### 基本用法

```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) GetCount() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

### 应用场景
- **共享数据**: 保护共享变量、map、slice等
- **资源管理**: 保护文件、网络连接等资源
- **状态同步**: 同步复杂状态变更

### 实际Demo

```go
func example_mutex_counter() {
    counter := &SafeCounter{}
    var wg sync.WaitGroup
    
    // 启动多个goroutine并发增加计数
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Final count: %d\n", counter.GetCount())
}
```

## Atomic

### 基本概念
Atomic包提供了原子操作，用于无锁的并发编程，性能比Mutex更好。

### 基本用法

```go
func example_atomic_counter() {
    var ops atomic.Int64
    var wg sync.WaitGroup
    
    // 启动多个goroutine
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                ops.Add(1)
            }
        }()
    }
    
    wg.Wait()
    fmt.Printf("ops: %d\n", ops.Load())
}
```

### 常用原子操作
- `Add`: 原子加法
- `Load`: 原子读取
- `Store`: 原子存储
- `CompareAndSwap`: 比较并交换
- `Swap`: 原子交换

### 应用场景
- **计数器**: 高性能计数器
- **标志位**: 状态标志
- **指针**: 原子指针操作

### 实际Demo

```go
type AtomicCounter struct {
    value atomic.Int64
}

func (c *AtomicCounter) Increment() {
    c.value.Add(1)
}

func (c *AtomicCounter) GetValue() int64 {
    return c.value.Load()
}

func (c *AtomicCounter) SetValue(val int64) {
    c.value.Store(val)
}
```

## Timer

### 基本概念
Timer用于在指定时间后执行一次操作。

### 基本用法

```go
func example_timer() {
    // 创建2秒定时器
    timer1 := time.NewTimer(time.Second * 2)
    
    // 等待定时器触发
    <-timer1.C
    fmt.Println("Timer 1 fired")
    
    // 停止定时器
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 fired")
    }()
    
    stop := timer2.Stop()
    if stop {
        fmt.Println("Timer 2 stopped")
    }
}
```

### 应用场景
- **超时控制**: 设置操作超时
- **延迟执行**: 延迟执行某个操作
- **定时清理**: 定期清理资源

### 实际Demo

```go
func example_timeout_operation() {
    ch := make(chan string)
    
    // 模拟耗时操作
    go func() {
        time.Sleep(time.Second * 3)
        ch <- "operation completed"
    }()
    
    // 设置超时
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-time.After(time.Second * 2):
        fmt.Println("operation timeout")
    }
}
```

## Ticker

### 基本概念
Ticker用于定期执行操作，类似于定时器但会重复触发。

### 基本用法

```go
func example_ticker() {
    ticker := time.NewTicker(time.Millisecond * 500)
    done := make(chan bool)
    
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
    
    time.Sleep(time.Second * 2)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
```

### 应用场景
- **定期任务**: 定期执行清理、备份等任务
- **心跳检测**: 定期发送心跳包
- **数据同步**: 定期同步数据
- **监控**: 定期检查系统状态

### 实际Demo

```go
func example_heartbeat() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    
    for i := 0; i < 5; i++ {
        select {
        case <-ticker.C:
            fmt.Printf("Heartbeat %d\n", i+1)
        }
    }
}
```

## 工具配合使用

### 1. Channel + WaitGroup + Goroutine

```go
func example_channel_waitgroup() {
    const numWorkers = 3
    const numJobs = 10
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    var wg sync.WaitGroup
    
    // 启动workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerId int) {
            defer wg.Done()
            for job := range jobs {
                fmt.Printf("Worker %d processing job %d\n", workerId, job)
                time.Sleep(time.Millisecond * 100)
                results <- job * 2
            }
        }(i)
    }
    
    // 发送任务
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 等待所有workers完成
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // 收集结果
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}
```

### 2. Mutex + Channel + Timer

```go
type Cache struct {
    mu    sync.RWMutex
    data  map[string]interface{}
    clean chan struct{}
}

func NewCache() *Cache {
    cache := &Cache{
        data:  make(map[string]interface{}),
        clean: make(chan struct{}),
    }
    
    // 定期清理过期数据
    go cache.cleanupRoutine()
    return cache
}

func (c *Cache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    value, exists := c.data[key]
    return value, exists
}

func (c *Cache) cleanupRoutine() {
    ticker := time.NewTicker(time.Minute)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            c.mu.Lock()
            // 清理逻辑
            c.mu.Unlock()
        case <-c.clean:
            return
        }
    }
}
```

### 3. Atomic + Channel + Select

```go
type RateLimiter struct {
    tokens    atomic.Int64
    maxTokens int64
    refill    time.Duration
    stop      chan struct{}
}

func NewRateLimiter(maxTokens int64, refill time.Duration) *RateLimiter {
    rl := &RateLimiter{
        maxTokens: maxTokens,
        refill:    refill,
        stop:      make(chan struct{}),
    }
    rl.tokens.Store(maxTokens)
    
    go rl.refillRoutine()
    return rl
}

func (rl *RateLimiter) Allow() bool {
    for {
        current := rl.tokens.Load()
        if current <= 0 {
            return false
        }
        
        if rl.tokens.CompareAndSwap(current, current-1) {
            return true
        }
    }
}

func (rl *RateLimiter) refillRoutine() {
    ticker := time.NewTicker(rl.refill)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            rl.tokens.Store(rl.maxTokens)
        case <-rl.stop:
            return
        }
    }
}
```

## 最佳实践

### 1. 避免Goroutine泄漏
```go
// 错误示例
func badExample() {
    go func() {
        // 无限循环，没有退出条件
        for {
            // 处理逻辑
        }
    }()
}

// 正确示例
func goodExample() {
    done := make(chan struct{})
    go func() {
        defer close(done)
        for {
            select {
            case <-done:
                return
            default:
                // 处理逻辑
            }
        }
    }()
    
    // 在适当时候关闭
    close(done)
}
```

### 2. 合理使用Channel缓冲
```go
// 无缓冲：同步通信
syncCh := make(chan string)

// 有缓冲：异步通信
asyncCh := make(chan string, 10)
```

### 3. 使用Context控制生命周期
```go
func exampleWithContext(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            default:
                // 处理逻辑
            }
        }
    }()
}
```

### 4. 避免死锁
```go
// 避免在同一个goroutine中发送和接收
func avoidDeadlock() {
    ch := make(chan string)
    
    // 错误：会死锁
    // ch <- "hello"
    // msg := <-ch
    
    // 正确：在不同goroutine中
    go func() {
        ch <- "hello"
    }()
    msg := <-ch
    fmt.Println(msg)
}
```

### 5. 性能优化
- 使用atomic代替mutex进行简单操作
- 合理使用channel缓冲
- 避免过度使用goroutine
- 使用sync.Pool复用对象

## 总结

Go的并发编程模型提供了丰富的工具来处理各种并发场景：

1. **Goroutine**: 轻量级线程，用于并发执行
2. **Channel**: 类型安全的通信机制
3. **WaitGroup**: 等待多个goroutine完成
4. **Mutex**: 保护共享资源
5. **Atomic**: 高性能原子操作
6. **Timer**: 单次定时操作
7. **Ticker**: 重复定时操作

这些工具可以灵活组合，构建出高效、安全的并发程序。关键是要理解每种工具的适用场景，并遵循Go的并发编程最佳实践。 