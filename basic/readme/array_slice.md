# Go 数组与切片深度解析

## 1. 核心概念理解

### 1.1 数组 (Array)

```go
// 数组是固定长度的值类型
// 数组由类型和长度定义
var arr1 [5]int                    // 零值数组
arr2 := [5]int{1, 2, 3, 4, 5}     // 字面量初始化
arr3 := [...]int{1, 2, 3}         // 编译器推导长度
arr4 := [5]int{1: 10, 3: 30}      // 指定索引初始化
```

### 1.2 切片 (Slice)

```go
// 切片是动态数组的抽象，引用类型
// 切片由类型定义
var s1 []int                       // 零值切片 (nil)
s2 := []int{1, 2, 3}              // 字面量初始化
s3 := make([]int, 5)              // 长度为5，容量为5
s4 := make([]int, 0, 10)          // 长度为0，容量为10
s5 := arr2[1:4]                   // 基于数组的切片
```

## 2. 内存布局与底层原理

### 2.1 数组内存布局

```go
// 数组在内存中是连续存储的
arr := [4]int{10, 20, 30, 40}
// 内存布局：[10][20][30][40]
// 地址：   addr addr+8 addr+16 addr+24 (假设int占8字节)
```

### 2.2 切片内存结构

```go
// 切片在运行时的结构体表示
type slice struct {
    array unsafe.Pointer // 指向底层数组的指针
    len   int            // 当前长度
    cap   int            // 容量
}
```

### 2.3 切片创建方式对比

```go
// 方式1: 字面量 - 创建数组后返回切片
s1 := []int{1, 2, 3}        // len=3, cap=3

// 方式2: make创建 - 直接分配内存
s2 := make([]int, 3, 5)     // len=3, cap=5，零值填充

// 方式3: 基于数组切片 - 共享底层数组
arr := [5]int{1, 2, 3, 4, 5}
s3 := arr[1:4]              // len=3, cap=4，共享arr的内存
```

## 3. 关键操作详解

### 3.1 切片扩容机制

```go
func demonstrateGrowth() {
    s := make([]int, 0, 1)
    fmt.Printf("初始: len=%d, cap=%d\n", len(s), cap(s))
    
    // 触发扩容
    for i := 0; i < 10; i++ {
        s = append(s, i)
        fmt.Printf("添加%d后: len=%d, cap=%d\n", i, len(s), cap(s))
    }
    
    // 扩容规则（简化版）：
    // - 当 cap < 1024 时，新容量 = 旧容量 * 2
    // - 当 cap >= 1024 时，新容量 = 旧容量 * 1.25
}
```

### 3.2 切片操作的陷阱

```go
// 陷阱1: 切片共享底层数组
func sliceShareArray() {
    arr := [5]int{1, 2, 3, 4, 5}
    s1 := arr[1:3]  // [2, 3]
    s2 := arr[2:4]  // [3, 4]
    
    s1[1] = 999     // 修改 s1 影响 s2
    fmt.Println(s2) // [999, 4] - 被意外修改！
}

// 陷阱2: 切片长度 vs 容量
func lengthVsCapacity() {
    s := make([]int, 0, 5)
    // s[0] = 1 // panic: 索引越界
    s = append(s, 1) // 正确方式
}

// 陷阱3: 切片复制
func sliceCopy() {
    s1 := []int{1, 2, 3}
    s2 := s1        // 浅复制，共享底层数组
    s2[0] = 999
    fmt.Println(s1) // [999, 2, 3] - s1 被修改
    
    // 深复制的正确方式
    s3 := make([]int, len(s1))
    copy(s3, s1)
    s3[0] = 111
    fmt.Println(s1) // [999, 2, 3] - s1 不受影响
}
```

## 4. 性能对比与最佳实践

### 4.1 性能对比

```go
func benchmarkArrayVsSlice() {
    // 数组传递 - 值拷贝，大数组性能差
    func processArray(arr [1000]int) {
        // 复制整个数组，内存开销大
    }
    
    // 切片传递 - 引用传递，性能好
    func processSlice(s []int) {
        // 只传递切片头，内存开销小
    }
}
```

### 4.2 最佳实践

#### 4.2.1 预分配容量

```go
// 不好的做法 - 频繁扩容
func badAppend() []int {
    var s []int
    for i := 0; i < 1000; i++ {
        s = append(s, i) // 可能触发多次扩容
    }
    return s
}

// 好的做法 - 预分配容量
func goodAppend() []int {
    s := make([]int, 0, 1000) // 预分配容量
    for i := 0; i < 1000; i++ {
        s = append(s, i) // 不会触发扩容
    }
    return s
}
```

#### 4.2.2 避免内存泄漏

```go
// 危险：保留大切片的小部分
func memoryLeak() []int {
    bigSlice := make([]int, 1000000)
    // ... 填充数据
    
    // 只需要前10个元素，但底层数组仍然很大
    return bigSlice[:10] // 内存泄漏！
}

// 安全：创建新切片
func avoidMemoryLeak() []int {
    bigSlice := make([]int, 1000000)
    // ... 填充数据
    
    // 创建新切片，释放大数组
    result := make([]int, 10)
    copy(result, bigSlice[:10])
    return result
}
```

#### 4.2.3 选择合适的数据结构

```go
// 使用数组的场景
func useArrayWhen() {
    // 1. 固定大小的数据
    var matrix [3][3]int
    
    // 2. 作为函数参数时需要值拷贝
    func process(data [32]byte) { /* ... */ }
    
    // 3. 编译时已知大小且不会变化
    var buffer [1024]byte
}

// 使用切片的场景
func useSliceWhen() {
    // 1. 动态大小的数据
    var items []string
    
    // 2. 作为函数参数避免拷贝
    func process(data []byte) { /* ... */ }
    
    // 3. 需要灵活操作的数据
    users := make([]User, 0, 100)
}
```

## 5. 实际应用场景

### 5.1 数据处理模式

```go
// 过滤模式
func filter(nums []int, predicate func(int) bool) []int {
    result := make([]int, 0, len(nums)) // 预分配
    for _, num := range nums {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

// 映射模式
func mapInt(nums []int, mapper func(int) int) []int {
    result := make([]int, len(nums)) // 预分配确定大小
    for i, num := range nums {
        result[i] = mapper(num)
    }
    return result
}

// 归约模式
func reduce(nums []int, initial int, reducer func(int, int) int) int {
    result := initial
    for _, num := range nums {
        result = reducer(result, num)
    }
    return result
}
```

### 5.2 缓冲区管理

```go
// 高效的缓冲区实现
type Buffer struct {
    buf []byte
    off int // 读取偏移
}

func (b *Buffer) Write(p []byte) (int, error) {
    // 扩容策略
    if len(b.buf) < b.off+len(p) {
        newBuf := make([]byte, 2*(b.off+len(p)))
        copy(newBuf, b.buf[:b.off])
        b.buf = newBuf
    }
    
    copy(b.buf[b.off:], p)
    b.off += len(p)
    return len(p), nil
}

func (b *Buffer) Read(p []byte) (int, error) {
    if b.off == 0 {
        return 0, io.EOF
    }
    
    n := copy(p, b.buf[:b.off])
    b.buf = b.buf[n:]
    b.off -= n
    return n, nil
}
```

# 6. 源码学习

```go
// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAll(r Reader) ([]byte, error) {
	b := make([]byte, 0, 512)  // len=0, cap=512
	for {
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]  // len=len(b)+n
		if err != nil {
			if err == EOF {
				err = nil
			}
			return b, err
		}

		if len(b) == cap(b) { // 底层数组满了
			// Add more capacity (let append pick how much).
            // b的容量满时，通过在新增加一个0元素，让append自动扩容底层数组长度，实现自动增加cap容量
            // 从而获取一个新的内存地址赋值给b，然后再把原来的有效元素赋值给新的数组，把0抛弃掉
			// 注意：这里只发生一次数组拷贝,append(b, 0)时，[:len(b)]只是改变新切片的长度。
            b = append(b, 0)[:len(b)]
		}
	}
}
```

应该深入理解切片各种操作的含义：

1. 切片的[begin: end]操作，实际只是修改len、cap这些量，并不会发生实际的数据复制
2. append在扩容时，会发生数组创建和复制

```go
a := []int{1, 2, 3, 4, 5} // len=5, cap=5
a = append(a, 6)          // len=6, cap=10
a = a[2:4]                // len=2, cap=8
fmt.Println(a)            // a=[3, 4]
```



# 7. 参考文档

https://i6448038.github.io/2018/08/11/array-and-slice-principle/

https://go.dev/blog/slices-intro#