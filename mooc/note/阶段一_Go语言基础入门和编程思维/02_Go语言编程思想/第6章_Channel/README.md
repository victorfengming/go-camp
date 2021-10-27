


# 6-1 channel

```go
package main

import "fmt"

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)
	c <- 1
	c <- 2
	c <- 3
	n:= <-c
	fmt.Println(n)

}

func main() {
	chanDemo()
}

/**
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.chanDemo()
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/channel.go:8 +0x37
main.main()
 */
```



> 死锁了

发的数据没人收是会deadlock的

![1635301183810](README/1635301183810.png)



# 6-2 使用Channel等待任务结束



# 6-3 使用Channel进行树的遍历



# 6-4 Select



# 6-5 传统同步机制



# 6-6 并发模式（上）



# 6-7 并发模式（下）



# 6-8 并发任务的控制