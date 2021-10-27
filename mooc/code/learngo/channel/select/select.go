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

func main() {
	var c1, c2 = generator(), generator()
	for {
		select {
		case n := <-c1:
			fmt.Println("received from c1:", n)
		case n := <-c2:
			fmt.Println("received from c2:", n)
			//default:
			//	fmt.Println("no val received")
		}
	}

}

/**
received from c1: 0
received from c2: 0
received from c2: 1
received from c1: 1
received from c1: 2
received from c2: 2
received from c1: 3
received from c2: 3
received from c2: 4
received from c2: 5
received from c1: 4
received from c1: 5
received from c1: 6
received from c2: 6
received from c1: 7
received from c2: 7
received from c1: 8

Process finished with the exit code -1073741510 (0xC000013A: interrupted by Ctrl+C)

*/
