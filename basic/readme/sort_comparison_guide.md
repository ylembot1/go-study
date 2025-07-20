# Go语言排序函数SortFunc和比较函数返回值关系详解

## 1. 核心概念

### 1.1 比较函数的返回值含义

```go
// 比较函数的签名
func(a, b T) int

// 返回值含义：
// -1: a < b  (a应该排在b前面)
//  0: a == b (a和b相等)
// +1: a > b  (a应该排在b后面)
```

### 1.2 SortFunc的工作原理

```go
// SortFunc使用比较函数来确定元素的相对顺序
// 它会调用比较函数来比较两个元素，然后根据返回值进行排序
slices.SortFunc(people, peopleCmp)
```

## 2. 详细示例分析

### 2.1 你的代码分析

```go
type peopleStruct struct {
    name string
    age  int
}

people := []peopleStruct{
    {name: "Jax", age: 37},
    {name: "Lena", age: 29},
    {name: "Mike", age: 32},
    {name: "Alice", age: 29},
    {name: "Bob", age: 32},
}

peopleCmp := func(a, b peopleStruct) int {
    if a.age == b.age {
        return cmp.Compare(a.name, b.name)  // 年龄相同时按姓名排序
    }
    return cmp.Compare(a.age, b.age)        // 按年龄排序
}

slices.SortFunc(people, peopleCmp)
```

### 2.2 排序过程演示

```go
func demonstrateSorting() {
    people := []peopleStruct{
        {name: "Jax", age: 37},
        {name: "Lena", age: 29},
        {name: "Mike", age: 32},
        {name: "Alice", age: 29},
        {name: "Bob", age: 32},
    }
    
    fmt.Println("排序前:", people)
    
    // 让我们手动模拟比较过程
    fmt.Println("\n比较过程:")
    
    // 比较 Lena(29) 和 Alice(29)
    // 年龄相同，比较姓名: "Lena" vs "Alice"
    // cmp.Compare("Lena", "Alice") 返回 +1 (Lena > Alice)
    // 所以 Alice 应该排在 Lena 前面
    
    // 比较 Mike(32) 和 Bob(32)  
    // 年龄相同，比较姓名: "Mike" vs "Bob"
    // cmp.Compare("Mike", "Bob") 返回 +1 (Mike > Bob)
    // 所以 Bob 应该排在 Mike 前面
    
    peopleCmp := func(a, b peopleStruct) int {
        if a.age == b.age {
            return cmp.Compare(a.name, b.name)
        }
        return cmp.Compare(a.age, b.age)
    }
    
    slices.SortFunc(people, peopleCmp)
    fmt.Println("排序后:", people)
    
    // 输出应该是：
    // [{Alice 29} {Lena 29} {Bob 32} {Mike 32} {Jax 37}]
}
```

## 3. 升序和降序控制

### 3.1 升序排序（默认）

```go
// 升序：从小到大
func ascendingSort() {
    numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
    
    // 升序比较函数
    ascendingCmp := func(a, b int) int {
        return cmp.Compare(a, b)  // 返回 a - b 的结果
    }
    
    slices.SortFunc(numbers, ascendingCmp)
    fmt.Println("升序:", numbers) // [1 1 2 3 4 5 6 9]
}
```

### 3.2 降序排序

```go
// 降序：从大到小
func descendingSort() {
    numbers := []int{3, 1, 4, 1, 5, 9, 2, 6}
    
    // 降序比较函数
    descendingCmp := func(a, b int) int {
        return cmp.Compare(b, a)  // 注意：交换了a和b的位置
    }
    
    slices.SortFunc(numbers, descendingCmp)
    fmt.Println("降序:", numbers) // [9 6 5 4 3 2 1 1]
}
```

### 3.3 多字段排序

```go
func multiFieldSort() {
    people := []peopleStruct{
        {name: "Alice", age: 30},
        {name: "Bob", age: 25},
        {name: "Alice", age: 25},
        {name: "Charlie", age: 30},
    }
    
    // 主排序：年龄升序，次排序：姓名升序
    primaryAscending := func(a, b peopleStruct) int {
        if a.age != b.age {
            return cmp.Compare(a.age, b.age)  // 年龄升序
        }
        return cmp.Compare(a.name, b.name)    // 姓名升序
    }
    
    // 主排序：年龄降序，次排序：姓名升序
    primaryDescending := func(a, b peopleStruct) int {
        if a.age != b.age {
            return cmp.Compare(b.age, a.age)  // 年龄降序
        }
        return cmp.Compare(a.name, b.name)    // 姓名升序
    }
    
    people1 := make([]peopleStruct, len(people))
    copy(people1, people)
    slices.SortFunc(people1, primaryAscending)
    fmt.Println("年龄升序，姓名升序:", people1)
    
    people2 := make([]peopleStruct, len(people))
    copy(people2, people)
    slices.SortFunc(people2, primaryDescending)
    fmt.Println("年龄降序，姓名升序:", people2)
}
```

## 4. 常见排序模式

### 4.1 字符串排序

```go
func stringSorting() {
    names := []string{"Charlie", "Alice", "Bob", "David"}
    
    // 升序（字典序）
    ascending := func(a, b string) int {
        return cmp.Compare(a, b)
    }
    
    // 降序（逆字典序）
    descending := func(a, b string) int {
        return cmp.Compare(b, a)
    }
    
    // 按长度排序
    byLength := func(a, b string) int {
        return cmp.Compare(len(a), len(b))
    }
    
    names1 := make([]string, len(names))
    copy(names1, names)
    slices.SortFunc(names1, ascending)
    fmt.Println("字典序升序:", names1)
    
    names2 := make([]string, len(names))
    copy(names2, names)
    slices.SortFunc(names2, descending)
    fmt.Println("字典序降序:", names2)
    
    names3 := make([]string, len(names))
    copy(names3, names)
    slices.SortFunc(names3, byLength)
    fmt.Println("按长度升序:", names3)
}
```

### 4.2 自定义类型排序

```go
type Product struct {
    Name     string
    Price    float64
    Category string
}

func customTypeSorting() {
    products := []Product{
        {Name: "Laptop", Price: 999.99, Category: "Electronics"},
        {Name: "Book", Price: 19.99, Category: "Books"},
        {Name: "Phone", Price: 699.99, Category: "Electronics"},
        {Name: "Pen", Price: 2.99, Category: "Office"},
    }
    
    // 按价格升序
    byPriceAsc := func(a, b Product) int {
        return cmp.Compare(a.Price, b.Price)
    }
    
    // 按价格降序
    byPriceDesc := func(a, b Product) int {
        return cmp.Compare(b.Price, a.Price)
    }
    
    // 按类别升序，然后按价格降序
    byCategoryThenPrice := func(a, b Product) int {
        if a.Category != b.Category {
            return cmp.Compare(a.Category, b.Category)
        }
        return cmp.Compare(b.Price, a.Price)  // 同类别内价格降序
    }
    
    products1 := make([]Product, len(products))
    copy(products1, products)
    slices.SortFunc(products1, byPriceAsc)
    fmt.Println("按价格升序:", products1)
    
    products2 := make([]Product, len(products))
    copy(products2, products)
    slices.SortFunc(products2, byCategoryThenPrice)
    fmt.Println("按类别升序，价格降序:", products2)
}
```

## 5. 理解cmp.Compare函数

### 5.1 cmp.Compare的工作原理

```go
// cmp.Compare(a, b) 返回：
// -1 如果 a < b
//  0 如果 a == b  
// +1 如果 a > b

func demonstrateCmpCompare() {
    fmt.Println("数字比较:")
    fmt.Printf("cmp.Compare(1, 2) = %d\n", cmp.Compare(1, 2))   // -1
    fmt.Printf("cmp.Compare(2, 2) = %d\n", cmp.Compare(2, 2))   //  0
    fmt.Printf("cmp.Compare(3, 2) = %d\n", cmp.Compare(3, 2))   // +1
    
    fmt.Println("\n字符串比较:")
    fmt.Printf("cmp.Compare(\"Alice\", \"Bob\") = %d\n", cmp.Compare("Alice", "Bob"))   // -1
    fmt.Printf("cmp.Compare(\"Bob\", \"Bob\") = %d\n", cmp.Compare("Bob", "Bob"))       //  0
    fmt.Printf("cmp.Compare(\"Charlie\", \"Bob\") = %d\n", cmp.Compare("Charlie", "Bob")) // +1
}
```

### 5.2 手动实现比较函数

```go
// 不使用cmp.Compare，手动实现比较逻辑
func manualComparison() {
    numbers := []int{3, 1, 4, 1, 5}
    
    // 手动实现升序比较
    ascendingManual := func(a, b int) int {
        if a < b {
            return -1
        } else if a > b {
            return 1
        }
        return 0
    }
    
    // 手动实现降序比较
    descendingManual := func(a, b int) int {
        if a > b {
            return -1
        } else if a < b {
            return 1
        }
        return 0
    }
    
    numbers1 := make([]int, len(numbers))
    copy(numbers1, numbers)
    slices.SortFunc(numbers1, ascendingManual)
    fmt.Println("手动升序:", numbers1)
    
    numbers2 := make([]int, len(numbers))
    copy(numbers2, numbers)
    slices.SortFunc(numbers2, descendingManual)
    fmt.Println("手动降序:", numbers2)
}
```

## 6. 实际应用场景

### 6.1 学生成绩排序

```go
type Student struct {
    Name   string
    Score  int
    Grade  string
}

func studentSorting() {
    students := []Student{
        {Name: "Alice", Score: 85, Grade: "B"},
        {Name: "Bob", Score: 92, Grade: "A"},
        {Name: "Charlie", Score: 78, Grade: "C"},
        {Name: "David", Score: 92, Grade: "A"},
    }
    
    // 按成绩降序，同成绩按姓名升序
    byScoreDescThenName := func(a, b Student) int {
        if a.Score != b.Score {
            return cmp.Compare(b.Score, a.Score)  // 成绩降序
        }
        return cmp.Compare(a.Name, b.Name)        // 姓名升序
    }
    
    slices.SortFunc(students, byScoreDescThenName)
    fmt.Println("按成绩降序，姓名升序:")
    for _, s := range students {
        fmt.Printf("  %s: %d (%s)\n", s.Name, s.Score, s.Grade)
    }
}
```

### 6.2 文件排序

```go
type FileInfo struct {
    Name    string
    Size    int64
    ModTime time.Time
}

func fileSorting() {
    files := []FileInfo{
        {Name: "document.txt", Size: 1024, ModTime: time.Now().Add(-24 * time.Hour)},
        {Name: "image.jpg", Size: 2048, ModTime: time.Now()},
        {Name: "video.mp4", Size: 1024, ModTime: time.Now().Add(-12 * time.Hour)},
    }
    
    // 按修改时间降序（最新的在前）
    byModTimeDesc := func(a, b FileInfo) int {
        return cmp.Compare(b.ModTime, a.ModTime)
    }
    
    // 按大小升序，同大小按名称升序
    bySizeAscThenName := func(a, b FileInfo) int {
        if a.Size != b.Size {
            return cmp.Compare(a.Size, b.Size)
        }
        return cmp.Compare(a.Name, b.Name)
    }
    
    files1 := make([]FileInfo, len(files))
    copy(files1, files)
    slices.SortFunc(files1, byModTimeDesc)
    fmt.Println("按修改时间降序:", files1)
    
    files2 := make([]FileInfo, len(files))
    copy(files2, files)
    slices.SortFunc(files2, bySizeAscThenName)
    fmt.Println("按大小升序，名称升序:", files2)
}
```

## 7. 总结

### 7.1 关键要点

1. **比较函数返回值**：
   - `-1`: a < b (a排在b前面)
   - `0`: a == b (a和b相等)
   - `+1`: a > b (a排在b后面)

2. **升序控制**：
   ```go
   return cmp.Compare(a, b)  // 升序
   ```

3. **降序控制**：
   ```go
   return cmp.Compare(b, a)  // 降序
   ```

4. **多字段排序**：
   ```go
   func(a, b T) int {
       if a.field1 != b.field1 {
           return cmp.Compare(a.field1, b.field1)  // 主字段
       }
       return cmp.Compare(a.field2, b.field2)      // 次字段
   }
   ```

### 7.2 记忆技巧

- **升序**：`cmp.Compare(a, b)` - 小的在前（升序），保持和函数传参一致
- **降序**：`cmp.Compare(b, a)` - 大的在前（降序），和函数传参相反
- **多字段**：先比较主字段，相等时再比较次字段

### 7.3 常见错误

```go
// 错误：忘记处理相等情况
func wrongComparison(a, b int) int {
    if a < b {
        return -1
    }
    return 1  // 错误！没有处理 a == b 的情况
}

// 正确：处理所有情况
func correctComparison(a, b int) int {
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}
```

理解这些关系后，你就能灵活控制排序的顺序了！ 