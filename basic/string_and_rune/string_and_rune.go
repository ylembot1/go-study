package main

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	const s = "你好，hello"

	fmt.Println("len(s):", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%d %d\n", idx, runeVal)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeVal, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d %d %d\n", i, runeVal, width)
		w = width

		examineRune(runeVal)
	}

	stringIteration()
}

func stringIteration() {
	s := "Hello, 世界"

	// 字节遍历
	fmt.Println("字节遍历:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("  索引%d: %d (%c)\n", i, s[i], s[i])
	}

	// 字符遍历 (使用range)
	fmt.Println("字符遍历:")
	for i, r := range s {
		fmt.Printf("  索引%d: %d (%c)\n", i, r, r)
	}

	stringProcessing()

	conversionCost()
}

func stringProcessing() {
	// 字符数统计
	countChars := func(s string) int {
		return len([]rune(s))
	}

	// 字符串反转
	reverseString := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	// 字符串截取
	substring := func(s string, start, end int) string {
		runes := []rune(s)
		if start < 0 || end > len(runes) || start > end {
			return ""
		}
		return string(runes[start:end])
	}

	// 测试
	text := "Hello, 世界!"
	fmt.Printf("原文: %s\n", text)
	fmt.Printf("字符数: %d\n", countChars(text))
	fmt.Printf("反转: %s\n", reverseString(text))
	fmt.Printf("截取(0,8): %s\n", []byte(substring(text, 0, 8)))
	fmt.Printf("截取(0,8): ", []byte(substring(text, 0, 8)))
	fmt.Println()
	fmt.Println([]byte("Hello, 世界!"))
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '你' {
		fmt.Println("found 你 sua")
	}
}

func conversionCost() {
	// string -> []byte: 24.125µs
	// string -> []rune: 190.813708ms
	// []byte -> string: 6.439625ms

	s := strings.Repeat("Hello, 世界", 1000)

	// 字符串到字节切片的转换
	start := time.Now()
	for i := 0; i < 10000; i++ {
		_ = []byte(s)
	}
	fmt.Printf("string -> []byte: %v\n", time.Since(start))

	// 字符串到符文切片的转换
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = []rune(s)
	}
	fmt.Printf("string -> []rune: %v\n", time.Since(start))

	// 字节切片到字符串的转换
	bytes := []byte(s)
	start = time.Now()
	for i := 0; i < 10000; i++ {
		_ = string(bytes)
	}
	fmt.Printf("[]byte -> string: %v\n", time.Since(start))
}
