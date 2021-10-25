package main

import (
	"fmt"
	"os"
	"strings"
	inn "week02/pkg/demo01/inn"
)

func main() {
	err := inn.A()
	if err != nil {
		fmt.Println(PrintMessage(err)) // 打印普通信息
		fmt.Println(PrintStack(err))   // 打印信息附带堆栈信息
		return
	}
	fmt.Println("ok")
}

// 打印普通信息，没有堆栈信息
func PrintMessage(err error) string {
	return err.Error()
}

// 打印详细信息，附带堆栈信息。
func PrintStack(err error) string {
	errMsg := fmt.Sprintf("%+v", err)
	return CleanPath(errMsg)
}

// 脱敏
func CleanPath(s string) string {
	return strings.ReplaceAll(s, GetCurrentPath()+"/", "")
}

// 获取当前项目目录
func GetCurrentPath() string {
	getwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return strings.Replace(getwd, "\\", "/", -1)
}
