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

// TODO 这里能够打印 出 完整的报错文件路径(堆栈信息)

/**
open abc: The system cannot find the file specified.
open abc: The system cannot find the file specified.
week02/pkg/demo01/inn.C
	pkg/demo01/inn/c.go:11
week02/pkg/demo01/inn.B
	pkg/demo01/inn/b.go:4
week02/pkg/demo01/inn.A
	pkg/demo01/inn/a.go:4
main.main
	pkg/demo01/main.go:11
runtime.main
	E:/env/Go/src/runtime/proc.go:255
runtime.goexit
	E:/env/Go/src/runtime/asm_amd64.s:1581

Process finished with the exit code 0

*/
