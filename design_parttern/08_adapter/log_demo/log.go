package main

import "fmt"

// --- 1. 目标接口 (Target Interface) ---
// 这是你的应用程序所期望的日志接口
type Logger interface {
	LogInfo(message string)
	LogError(message string)
}

// --- 2. 被适配者 (Adaptee) ---
// 这是一个第三方日志库，它的接口与我们的 Logger 不兼容
// 我们无法（或不想）修改它的代码
type AwsCloudWatchLogger struct {
	// 可能有一些 AWS 相关的配置字段
}

// 它的方法名和签名都与我们的 Logger 不同
func (a *AwsCloudWatchLogger) SendLog(level string, msg string) {
	// 模拟将日志发送到 AWS CloudWatch
	fmt.Printf("[AWS CloudWatch - %s]: %s\n", level, msg)
}

// --- 3. 适配器 (Adapter) ---
// 适配器结构体：
// 1. 需要实现目标接口 (Logger)
// 2. 内部持有一个被适配者的实例 (*AwsCloudWatchLogger)
type CloudWatchAdapter struct {
	awsLogger *AwsCloudWatchLogger
}

// 工厂函数，用于创建适配器
func NewCloudWatchAdapter(awsLogger *AwsCloudWatchLogger) *CloudWatchAdapter {
	return &CloudWatchAdapter{awsLogger: awsLogger}
}

// 实现目标接口的方法，在内部进行“转换”
func (a *CloudWatchAdapter) LogInfo(message string) {
	// 将我们系统的 LogInfo 调用“适配”成第三方库的 SendLog 调用
	a.awsLogger.SendLog("INFO", message)
}

func (a *CloudWatchAdapter) LogError(message string) {
	// 将我们系统的 LogError 调用“适配”成第三方库的 SendLog 调用
	a.awsLogger.SendLog("ERROR", message)
}

// --- 4. 客户端代码 ---
// 我们的应用程序业务逻辑，它只认识 Logger 接口
func applicationLogic(logger Logger) {
	fmt.Println("Application is running...")
	logger.LogInfo("User logged in")
	logger.LogError("Failed to connect to database")
}

func main() {
	fmt.Println("=== Using the AWS CloudWatch Logger via Adapter ===")
	// 创建第三方日志实例
	awsLogger := &AwsCloudWatchLogger{}
	// 用适配器将它包装起来，使其符合 Logger 接口
	adapter := NewCloudWatchAdapter(awsLogger)
	// 将适配器传入业务逻辑，业务逻辑对此无感知
	applicationLogic(adapter)
}
