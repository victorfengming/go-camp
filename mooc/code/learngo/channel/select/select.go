package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) * time.Millisecond,
			)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for {
		n, ok := <-c
		if ok {
			fmt.Printf("@%d---%d\n", id, n)
		} else {
			break
		}
	}
}

// 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int) chan int { // 告诉外面用的人 , 我这个channel怎么用
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var work = createWorker(0)

	n := 0
	hasValue := false
	for {

		var activeWorker chan<- int
		if hasValue {
			activeWorker = work
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}

}

/**
@0---0
@0---0
@0---1
@0---1
@0---2
@0---2
@0---3

Process finished with the exit code -1073741510 (0xC000013A: interrupted by Ctrl+C)

*/
