# 基础-常用方法

```go
// sort
nums := []int{5, 4, 3, 1, 2}
sort.Ints(nums) // {1, 2, 3, 4, 5}

words := []string{"banana", "apple", "cherry"}
sort.Strings(words)  // [apple banana cherry]
```


```go
// slices
nums := []int{5, 4, 3, 1, 2}
m := slices.Min(nums)  // 1
M := slices.Max(nums)  // 5
```