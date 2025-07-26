package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// 顶层结构体
type Config struct {
	Title    string            `toml:"title"`
	Server   ServerConfig      `toml:"server"`
	Database DatabaseConfig    `toml:"database"`
	Users    []User            `toml:"users"` // TOML 数组用切片表示
	Map      map[string]string `toml:"map"`
}

// 嵌套结构体
type ServerConfig struct {
	Port    int  `toml:"port"`
	Enabled bool `toml:"enabled"`
}

type DatabaseConfig struct {
	Host        string `toml:"host"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	Connections int    `toml:"connections"`
}

type User struct {
	Name  string   `toml:"name"`
	Roles []string `toml:"roles"`
}

func LoadConf(fileName string, data interface{}) error {
	fileName, err := filepath.Abs(fileName)
	if err != nil {
		return err
	}

	bs, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return toml.Unmarshal(bs, data)
}

func main() {
	var config Config

	// 解析 TOML 文件
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	// 打印解析结果
	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Server Port: %d, Enabled: %v\n", config.Server.Port, config.Server.Enabled)
	fmt.Printf("DB Host: %s, User: %s\n", config.Database.Host, config.Database.Username)

	for i, user := range config.Users {
		fmt.Printf("User %d: %s, Roles: %v\n", i+1, user.Name, user.Roles)
	}

	fmt.Printf("Map: %v\n", config.Map)

	var config2 Config
	LoadConf("config.toml", &config2)
	fmt.Printf("Title: %s\n", config2.Title)
	fmt.Printf("Server Port: %d, Enabled: %v\n", config2.Server.Port, config2.Server.Enabled)
	fmt.Printf("DB Host: %s, User: %s\n", config2.Database.Host, config2.Database.Username)

	for i, user := range config2.Users {
		fmt.Printf("User %d: %s, Roles: %v\n", i+1, user.Name, user.Roles)
	}

	fmt.Printf("Map: %v\n", config2.Map)
}
