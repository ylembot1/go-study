package main

import (
	"fmt"
	"strings"
)

// --- 目标接口 ---
type JsonDataProvider interface {
	GetJsonData() string
}

// --- 被适配者：一个返回XML的旧服务 ---
type OldXmlService struct{}

func (s *OldXmlService) GetXmlData() string {
	// 模拟返回XML数据
	return `<user><name>Alice</name><email>alice@example.com</email></user>`
}

// --- 适配器：将XML服务适配成JSON提供商 ---
type XmlToJsonAdapter struct {
	xmlService *OldXmlService
}

func NewXmlToJsonAdapter(service *OldXmlService) *XmlToJsonAdapter {
	return &XmlToJsonAdapter{xmlService: service}
}

func (a *XmlToJsonAdapter) GetJsonData() string {
	// 1. 从旧服务获取XML数据
	xmlData := a.xmlService.GetXmlData()

	// 2. 这里进行一个非常非常简单（且不严谨）的XML到JSON的转换演示
	// 真正的转换需要使用正式的库，这里只是为了展示适配器的思想
	jsonData := strings.NewReplacer(
		"<user>", `{"user": {`,
		"<name>", `"name": "`,
		"</name>", `",`,
		"<email>", `"email": "`,
		"</email>", `"`,
		"</user>", `}}`,
	).Replace(xmlData)

	fmt.Printf("Adapter: Converted XML '%s' to JSON '%s'\n", xmlData, jsonData)
	return jsonData
}

// --- 客户端代码 ---
// 新系统的一个组件，它只接受JSON数据
func NewSystemComponent(provider JsonDataProvider) {
	fmt.Println("New system component received JSON data:", provider.GetJsonData())
}

func main() {
	fmt.Println("=== Integrating the Old XML Service into the New JSON System ===")

	oldService := &OldXmlService{}
	adapter := NewXmlToJsonAdapter(oldService)

	// 现在可以把适配器当作一个 JsonDataProvider 来使用了
	NewSystemComponent(adapter)
}
