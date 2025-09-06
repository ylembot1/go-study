package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var num = "2799"
var folderName string = "count-complete-subarrays-in-an-array"

func Test_init(t *testing.T) {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前目录失败:", err)
		return
	}

	dirPath := filepath.Join(currentDir, num+"_"+folderName)
	if err := os.Mkdir(dirPath, 0755); err != nil {
		if os.IsExist(err) {
			fmt.Printf("文件夹 %s 已存在\n", dirPath)
		} else {
			fmt.Println("创建文件夹失败:", err)
		}
		return
	}

	filePath := filepath.Join(dirPath, folderName)
	file, err := os.Create(filePath + ".go")
	fileTest, err := os.Create(filePath + "_test.go")
	_, err = file.WriteString("package leetcode\n\n")
	_, err = fileTest.WriteString("package leetcode\n\nfunc Test_(t *testing.T) {\n\n}")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()
	defer fileTest.Close()
	fmt.Printf("成功创建文件 %s\n", filePath)

	fmt.Println("hello")
}
