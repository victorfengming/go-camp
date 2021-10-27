package main

import (
	"fmt"
	"math/rand"
	"time"
)

// chan bool 或者 chan struct{}
// 这个 chan struct{} 里面没有任何的数据,他比bool更加的省空间
func msgGen(name string, done chan struct{}) <-chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			// 判断 done 是否在
			select {
			case <-time.After(time.Millisecond * time.Duration(rand.Intn(5000))):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				// 证明我是主动的退出
				fmt.Println("cleaning up")
				time.Sleep(2 * time.Second)
				fmt.Println("cleaning over")
				//done <- "cleaning over"
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

func fanIn(chs ...<-chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		//chTemp := ch
		go func(in <-chan string) {
			// todo c1,c2 如何调度
			for {
				c <- <-in
			}
		}(ch)
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

func nonBlockingWait(c <-chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

func timeoutWait(c <-chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		// 没有等到
		return "", false
	}
}

func main() {

	done := make(chan struct{})
	// 生成消息
	m1 := msgGen("服务A", done)
	m2 := msgGen("服务B", done)

	for i := 0; i < 5; i++ {
		fmt.Println(<-m1)
		if m, ok := timeoutWait(m2, 1*time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("no message from serve")
			//time.Sleep(2*time.Second)
		}
	}
	done <- struct{}{}
	// 送完停止信号,为了让看到结果,再等他个2s钟
	<-done

}

/**
service 服务A: message 0
service 服务B: message 0
service 服务A: message 1
no message from serve
service 服务A: message 2
service 服务B: message 1
service 服务A: message 3
no message from serve
service 服务A: message 4
service 服务B: message 2
cleaning up
cleaning over

Process finished with the exit code 0


*/
