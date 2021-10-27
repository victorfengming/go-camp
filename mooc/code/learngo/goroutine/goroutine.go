package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//fmt.Printf("hellow from goroutine %d")
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			// race condition
			for {
				a[i]++
				//fmt.Printf("hellow from goroutine %d\n", ii)
				// 交出控制权
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

/**
PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\goroutine> go run .\goroutine.go -race
panic: runtime error: index out of range [10] with length 10

goroutine 15 [running]:
main.main.func1()
        E:/Projects/GolandProjects/go-camp/mooc/code/learngo/goroutine/goroutine.go:16 +0x56
created by main.main
        E:/Projects/GolandProjects/go-camp/mooc/code/learngo/goroutine/goroutine.go:13 +0x57
panic: runtime error: index out of range [10] with length 10

goroutine 10 [running]:
main.main.func1()
        E:/Projects/GolandProjects/go-camp/mooc/code/learngo/goroutine/goroutine.go:16 +0x56
created by main.main
        E:/Projects/GolandProjects/go-camp/mooc/code/learngo/goroutine/goroutine.go:13 +0x57
exit status 2
PS E:\Projects\GolandProjects\go-camp\mooc\code\learngo\goroutine>

*/
