package main

import "fmt"

// 1. 核心接口
type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// 2. 具体的组件，实现核心功能（被装饰的对象）
type FileDataSource struct {
	data string
}

func (f *FileDataSource) WriteData(data string) {
	fmt.Printf("Writing '%s' to the file.\n", data)
	f.data = data
}

func (f *FileDataSource) ReadData() string {
	fmt.Printf("Reading '%s' from the file.\n", f.data)
	return f.data
}

// 3. 基础装饰器结构体（通常是一个抽象层，非必须，但能让具体装饰器更简单）
type DataSourceDecorator struct {
	wrappee DataSource // 关键：持有被装饰的接口实例
}

// 装饰器本身也实现相同的接口，默认调用被装饰对象的方法
func (d *DataSourceDecorator) WriteData(data string) {
	d.wrappee.WriteData(data)
}

func (d *DataSourceDecorator) ReadData() string {
	return d.wrappee.ReadData()
}

// 4. 具体的装饰器A：加密装饰器
type EncryptionDecorator struct {
	DataSourceDecorator // 嵌入基础装饰器，继承其 wrappee 和默认方法
}

// 工厂函数，用于创建加密装饰器
func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{
		DataSourceDecorator: DataSourceDecorator{wrappee: source},
	}
}

// 重写接口方法，添加新功能
func (e *EncryptionDecorator) WriteData(data string) {
	encryptedData := "encrypted(" + data + ")" // 模拟加密过程
	fmt.Printf("EncryptionDecorator: Encrypting data to '%s'\n", encryptedData)
	// 调用被装饰对象的原始方法，但传入的是处理后的数据
	e.wrappee.WriteData(encryptedData)
}

func (e *EncryptionDecorator) ReadData() string {
	originalData := e.wrappee.ReadData()
	// 读取到加密后的数据，需要解密
	decryptedData := originalData[len("encrypted(") : len(originalData)-1] // 模拟解密
	fmt.Printf("EncryptionDecorator: Decrypting data to '%s'\n", decryptedData)
	return decryptedData
}

// 5. 具体的装饰器B：压缩装饰器
type CompressionDecorator struct {
	DataSourceDecorator
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{
		DataSourceDecorator: DataSourceDecorator{wrappee: source},
	}
}

func (c *CompressionDecorator) WriteData(data string) {
	compressedData := "compressed(" + data + ")" // 模拟压缩
	fmt.Printf("CompressionDecorator: Compressing data to '%s'\n", compressedData)
	c.wrappee.WriteData(compressedData)
}

func (c *CompressionDecorator) ReadData() string {
	originalData := c.wrappee.ReadData()
	decompressedData := originalData[len("compressed(") : len(originalData)-1] // 模拟解压
	fmt.Printf("CompressionDecorator: Decompressing data to '%s'\n", decompressedData)
	return decompressedData
}

// 6. 客户端使用
func main() {
	fmt.Println("=== 1. Basic Component ===")
	var source DataSource = &FileDataSource{}
	source.WriteData("My Design Pattern")
	fmt.Println(source.ReadData())

	fmt.Println("\n=== 2. Decorated with Encryption ===")
	var encryptedSource DataSource = NewEncryptionDecorator(source)
	encryptedSource.WriteData("My Secret Data")
	fmt.Println(encryptedSource.ReadData())

	fmt.Println("\n=== 3. Decorated with Compression AND Encryption ===")
	// 装饰器可以嵌套！顺序很重要。
	// 先压缩，再加密
	var superSource DataSource = NewEncryptionDecorator(
		NewCompressionDecorator(source),
	)
	superSource.WriteData("Very Large and Sensitive Data")
	fmt.Println(superSource.ReadData())
}
