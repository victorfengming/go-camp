package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func chanDemo() {
	//var c chan int // c == nil
	c := make(chan int)

	go worker(c)

	c <- 1
	c <- 2
	//c <- 3
	//n := <-c
	//fmt.Println(n)
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
1
2

Process finished with the exit code 0
*/
