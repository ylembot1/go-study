# Go 语言中 new 和 make 的区别与用法

## 1. 基础概念

### 1.1 new 关键字
```go
// new 是一个内置函数，用于分配内存
func new(Type) *Type
```
- **作用**: 分配内存并返回指向该内存的指针
- **返回值**: 指向类型零值的指针
- **适用于**: 所有类型

### 1.2 make 关键字
```go
// make 是一个内置函数，用于创建和初始化特定类型
func make(t Type, size ...IntegerType) Type
```
- **作用**: 创建、初始化并返回类型本身
- **返回值**: 类型本身（不是指针）
- **适用于**: 仅限于 slice、map、channel 三种类型

## 2. 核心区别对比

| 特性 | new | make |
|------|-----|------|
| **返回类型** | 指针 (*T) | 类型本身 (T) |
| **内存初始化** | 零值 | 已初始化，可用 |
| **适用类型** | 所有类型 | slice、map、channel |
| **是否可用** | 需要解引用 | 直接可用 |

## 3. 详细用法示例

### 3.1 基本类型的使用

```go
package main

import "fmt"

func basicTypeExample() {
    // 使用 new 创建基本类型
    var p1 *int = new(int)
    fmt.Printf("new(int): %T, 值: %v, 指向: %v\n", p1, p1, *p1)
    
    // 等价的传统方式
    var i int
    var p2 *int = &i
    fmt.Printf("&variable: %T, 值: %v, 指向: %v\n", p2, p2, *p2)
    
    // make 不能用于基本类型
    // var p3 = make(int) // 编译错误！
}
```

### 3.2 结构体的使用

```go
type Person struct {
    Name string
    Age  int
}

func structExample() {
    // 使用 new 创建结构体
    p1 := new(Person)
    fmt.Printf("new(Person): %T, 值: %v\n", p1, p1)
    fmt.Printf("结构体内容: %+v\n", *p1) // {Name: Age:0}
    
    // 需要解引用才能使用
    (*p1).Name = "Alice"
    p1.Age = 30 // Go 自动解引用
    
    // 等价的传统方式
    p2 := &Person{}
    p2.Name = "Bob"
    p2.Age = 25
    
    // make 不能用于结构体
    // p3 := make(Person) // 编译错误！
}
```

### 3.3 切片的使用

```go
func sliceExample() {
    // 使用 new 创建切片 - 返回指向 nil 切片的指针
    s1 := new([]int)
    fmt.Printf("new([]int): %T, 值: %v\n", s1, s1)
    fmt.Printf("切片内容: %v, 长度: %d, 容量: %d\n", *s1, len(*s1), cap(*s1))
    
    // 需要解引用并初始化才能使用
    *s1 = append(*s1, 1, 2, 3)
    
    // 使用 make 创建切片 - 返回可用的切片
    s2 := make([]int, 3, 5) // 长度3，容量5
    fmt.Printf("make([]int, 3, 5): %T, 值: %v\n", s2, s2)
    fmt.Printf("切片内容: %v, 长度: %d, 容量: %d\n", s2, len(s2), cap(s2))
    
    // 直接可用
    s2[0] = 10
    s2[1] = 20
    s2[2] = 30
    
    // 不同的创建方式
    s3 := make([]int, 0, 10)    // 长度0，容量10
    s4 := make([]int, 5)        // 长度5，容量5
    s5 := []int{1, 2, 3}        // 字面量创建
}
```

### 3.4 映射的使用

```go
func mapExample() {
    // 使用 new 创建映射 - 返回指向 nil map 的指针
    m1 := new(map[string]int)
    fmt.Printf("new(map[string]int): %T, 值: %v\n", m1, m1)
    fmt.Printf("映射内容: %v\n", *m1) // map[]
    
    // 需要初始化才能使用
    *m1 = make(map[string]int)
    (*m1)["key1"] = 100
    
    // 使用 make 创建映射 - 返回可用的映射
    m2 := make(map[string]int)
    fmt.Printf("make(map[string]int): %T, 值: %v\n", m2, m2)
    
    // 直接可用
    m2["key1"] = 100
    m2["key2"] = 200
    
    // 带初始容量的创建
    m3 := make(map[string]int, 10) // 提示初始容量
}
```

### 3.5 通道的使用

```go
func channelExample() {
    // 使用 new 创建通道 - 返回指向 nil channel 的指针
    ch1 := new(chan int)
    fmt.Printf("new(chan int): %T, 值: %v\n", ch1, ch1)
    fmt.Printf("通道内容: %v\n", *ch1) // <nil>
    
    // 需要初始化才能使用
    *ch1 = make(chan int)
    
    // 使用 make 创建通道 - 返回可用的通道
    ch2 := make(chan int)        // 无缓冲通道
    ch3 := make(chan int, 5)     // 缓冲通道，容量5
    
    fmt.Printf("make(chan int): %T, 值: %v\n", ch2, ch2)
    fmt.Printf("make(chan int, 5): %T, 值: %v\n", ch3, ch3)
    
    // 直接可用
    go func() {
        ch2 <- 42
        ch3 <- 100
    }()
}
```

## 4. 常见错误和陷阱

### 4.1 使用 new 创建引用类型的陷阱

```go
func commonMistakes() {
    // 错误1：使用 new 创建切片但忘记初始化
    s1 := new([]int)
    // s1[0] = 1 // panic: runtime error: index out of range
    
    // 正确做法1：使用 make
    s2 := make([]int, 5)
    s2[0] = 1 // 正常工作
    
    // 正确做法2：使用 new 后初始化
    s3 := new([]int)
    *s3 = make([]int, 5)
    (*s3)[0] = 1 // 正常工作
    
    // 错误2：使用 new 创建映射但忘记初始化
    m1 := new(map[string]int)
    // (*m1)["key"] = 1 // panic: assignment to entry in nil map
    
    // 正确做法
    m2 := make(map[string]int)
    m2["key"] = 1 // 正常工作
}
```

### 4.2 零值和初始化的区别

```go
func zeroValueVsInitialized() {
    // new 创建零值
    var s1 []int        // nil 切片
    s2 := new([]int)    // 指向 nil 切片的指针
    
    fmt.Printf("var s1: %v, nil? %v\n", s1, s1 == nil)           // true
    fmt.Printf("*new([]int): %v, nil? %v\n", *s2, *s2 == nil)    // true
    
    // make 创建初始化的值
    s3 := make([]int, 0)    // 空切片，但不是 nil
    s4 := make([]int, 3)    // 长度3的切片，元素为零值
    
    fmt.Printf("make([]int, 0): %v, nil? %v\n", s3, s3 == nil)   // false
    fmt.Printf("make([]int, 3): %v, nil? %v\n", s4, s4 == nil)   // false
}
```

## 5. 使用场景和最佳实践

### 5.1 何时使用 new

```go
// 1. 创建基本类型的指针
func useNewWhen() {
    // 需要指向基本类型的指针
    counter := new(int)
    *counter = 42
    
    // 创建结构体指针
    person := new(Person)
    person.Name = "Alice"
    
    // 在函数中需要返回指针
    func createCounter() *int {
        return new(int) // 返回指向零值的指针
    }
}

// 2. 泛型编程中的使用
func generic[T any]() *T {
    return new(T) // 创建类型 T 的零值指针
}
```

### 5.2 何时使用 make

```go
// 1. 创建引用类型时总是使用 make
func useMakeWhen() {
    // 创建切片
    numbers := make([]int, 0, 100)    // 预分配容量
    
    // 创建映射
    cache := make(map[string]interface{})
    
    // 创建通道
    jobs := make(chan Job, 10)        // 缓冲通道
    
    // 根据参数动态创建
    func createSlice(size int) []int {
        return make([]int, size)
    }
}
```

### 5.3 性能考虑

```go
func performanceConsiderations() {
    // 预分配容量避免频繁扩容
    items := make([]Item, 0, 1000)    // 好的做法
    
    // 而不是
    var items2 []Item                 // 可能需要多次扩容
    
    // 映射的预分配
    cache := make(map[string]string, 100)  // 提示初始容量
    
    // 通道的缓冲
    tasks := make(chan Task, 50)      // 减少协程阻塞
}
```

## 6. 内存分配对比

### 6.1 内存分配位置

```go
func memoryAllocation() {
    // new 分配在堆上（如果逃逸分析决定）
    p := new(int)
    *p = 42
    
    // make 也分配在堆上
    s := make([]int, 100)
    
    // 局部变量可能在栈上
    var local int = 42
    
    // 取地址可能导致逃逸到堆
    ptr := &local
    fmt.Println(ptr)
}
```

### 6.2 垃圾回收影响

```go
func gcImpact() {
    // new 创建的对象需要垃圾回收
    for i := 0; i < 1000; i++ {
        p := new(LargeStruct)
        // 使用 p...
        // p 会被垃圾回收器回收
    }
    
    // 重用对象减少 GC 压力
    var pool = sync.Pool{
        New: func() interface{} {
            return new(LargeStruct)
        },
    }
    
    obj := pool.Get().(*LargeStruct)
    // 使用 obj...
    pool.Put(obj) // 重用对象
}
```

## 7. 实际应用示例

### 7.1 工厂模式

```go
type Config struct {
    Host string
    Port int
}

// 使用 new 的工厂函数
func NewConfig() *Config {
    return new(Config) // 返回指向零值的指针
}

// 使用字面量的工厂函数
func NewConfigWithDefaults() *Config {
    return &Config{
        Host: "localhost",
        Port: 8080,
    }
}

// 使用 make 创建复杂类型
func NewCache() map[string]interface{} {
    return make(map[string]interface{})
}
```

### 7.2 数据结构初始化

```go
type DataProcessor struct {
    buffer []byte
    cache  map[string]interface{}
    done   chan bool
}

func NewDataProcessor(bufferSize int) *DataProcessor {
    return &DataProcessor{
        buffer: make([]byte, bufferSize),        // 使用 make
        cache:  make(map[string]interface{}),    // 使用 make
        done:   make(chan bool),                 // 使用 make
    }
}
```

## 8. 总结

### 8.1 快速决策指南

```
需要创建什么？
├── 基本类型（int, string, bool 等）
│   └── 使用 new 或 &variable
├── 结构体
│   └── 使用 new 或 &struct{}
├── 切片
│   └── 使用 make
├── 映射
│   └── 使用 make
└── 通道
    └── 使用 make
```

### 8.2 记忆口诀

- **new**: 分配内存，返回指针，适用所有类型
- **make**: 创建初始化，返回类型，仅限三种类型（slice、map、channel）

### 8.3 最佳实践

1. **引用类型优先使用 make**：切片、映射、通道
2. **需要指针时使用 new**：基本类型、结构体
3. **预分配容量**：避免频繁扩容
4. **理解零值**：new 返回零值，make 返回可用值
5. **性能考虑**：合理选择减少内存分配和 GC 压力

记住：**new 分配，make 创建**！ 