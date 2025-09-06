package main

import "fmt"

// 给你一个二进制字符串 s ，该字符串 不含前导零 。

// 如果 s 包含 零个或一个由连续的 '1' 组成的字段 ，返回 true​​​ 。否则，返回 false 。

func checkOnesSegment(s string) bool {
	has_one := false
	has_zero := false

	for _, char := range s {
		if char == '0' {
			if has_one {
				has_zero = true
			}
		} else {
			if has_zero && has_one {
				return false
			}
			has_one = true
		}
	}

	return true
}

func main() {
	fmt.Println(checkOnesSegment("110"))  // true - 只有一个连续的'1'字段
	fmt.Println(checkOnesSegment("1001")) // false - 有两个'1'字段
	fmt.Println(checkOnesSegment("000"))  // true - 零个'1'字段
	fmt.Println(checkOnesSegment("111"))  // true - 一个连续的'1'字段
	fmt.Println(checkOnesSegment("1"))    // true - 一个连续的'1'字段
}
