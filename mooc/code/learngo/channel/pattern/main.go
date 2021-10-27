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

func fanIn(chs ...<-chan string) chan string {
	c := make(chan string)
	var ch_temp <-chan string
	for _, ch := range chs {
		ch_temp = ch
		go func() {
			// todo c1,c2 如何调度
			for {
				c <- <-ch_temp
			}
		}()
	}

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
	m3 := msgGen("服务C")
	m := fanIn(m1, m2, m3)

	for {
		fmt.Println(<-m)
		//m<- "abc"
		// 没有办法发数据
	}

}

/**
service 服务C: message 0
service 服务C: message 1
service 服务C: message 2
service 服务C: message 3
service 服务C: message 4
service 服务C: message 5
service 服务C: message 6

Process finished with the exit code -1073741510 (0xC000013A: interrupted by Ctrl+C)

*/
