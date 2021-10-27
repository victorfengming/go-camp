package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("@%d---%c\n", id, <-c)
	}
}

// // 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int) chan<- int { // 告诉外面用的人 , 我这个channel怎么用
	//
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("@%d---%c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {

	var cahnneles [10]chan<- int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = createWorker(i)
		//n:= <-cahnneles[i]
		// ↑ Invalid operation: <-cahnneles[i] (receive from the send-only type chan<- int)
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

}

func bufferedChannel() {
	// 加上缓冲区,大小为3
	c := make(chan int, 3)
	go worker(0, c)
	c <- '1'
	c <- '2'
	c <- '3'
	//c <- 4
	time.Sleep(time.Millisecond)

}

func main() {
	//chanDemo()
	bufferedChannel()
}

/**
@0---1
@0---2
@0---3

Process finished with the exit code 0

*/
