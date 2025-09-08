package main

import "fmt"

// HttpHandler 是核心函数类型，类似于接口
type HttpHandler func(string)

// 原始处理函数
func myHandler(msg string) {
	fmt.Printf("Original Handler: Processing request - %s\n", msg)
}

// 装饰器函数：接收一个 Handler，返回一个增强后的 Handler
func loggingDecorator(h HttpHandler) HttpHandler {
	// 返回一个闭包，这个闭包就是新的 Handler
	return func(msg string) {
		fmt.Printf("LOG: Started handling request - %s\n", msg)
		h(msg) // 调用被装饰的原始函数
		fmt.Printf("LOG: Finished handling request - %s\n", msg)
	}
}

func authDecorator(h HttpHandler) HttpHandler {
	return func(msg string) {
		fmt.Println("AUTH: Checking credentials... [OK]")
		h(msg)
	}
}

func main() {
	fmt.Println("--- Original Handler ---")
	myHandler("Hello")

	fmt.Println("\n--- With Logging ---")
	decoratedHandler := loggingDecorator(myHandler)
	decoratedHandler("Hello")

	fmt.Println("\n--- With Auth AND Logging ---")
	// 装饰器可以链式调用
	superHandler := authDecorator(loggingDecorator(myHandler))
	superHandler("Super Request")
}
