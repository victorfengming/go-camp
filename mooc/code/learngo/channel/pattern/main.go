package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan 是一等公民
func msgGen() <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
			// sprintf 作为字符串打印
			c <- fmt.Sprintf("message %d", i)
			i++
		}
	}()
	return c
}
func main() {
	// 生成消息
	m1 := msgGen()
	m2 := msgGen()
	for {
		fmt.Println(<-m1)
		fmt.Println(<-m2)
		//m<- "abc"
		// 没有办法发数据
	}

}

/**
message 0
message 0
message 1
message 1
message 2
message 2
message 3
message 3
message 4
message 4
message 5
message 5
message 6
message 6
message 7


*/
