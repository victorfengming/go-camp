package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan 是一等公民
func msgGen(name string) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			// sprintf 作为字符串打印
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func fanIn(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		// todo c1,c2 如何调度
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

func fanInBySelect(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	// 生成消息
	m1 := msgGen("服务A")
	m2 := msgGen("服务B")
	m := fanInBySelect(m1, m2)

	for {
		fmt.Println(<-m)
		//m<- "abc"
		// 没有办法发数据
	}

}

/**
service 服务A: message 0
service 服务B: message 0
service 服务A: message 1
service 服务B: message 1
service 服务A: message 2
service 服务A: message 3
service 服务A: message 4
service 服务B: message 2
service 服务A: message 5
service 服务A: message 6
service 服务B: message 3


Process finished with the exit code -1073741510 (0xC000013A: interrupted by Ctrl+C)

*/
