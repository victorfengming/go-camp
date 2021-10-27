package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("%d,%c\n", id, <-c)
	}
}

func chanDemo() {

	var cahnneles [10]chan int
	for i := 0; i < 10; i++ {
		//var c chan int // c == nil
		cahnneles[i] = make(chan int)
		go worker(i, cahnneles[i])
	}

	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		cahnneles[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)

}

func main() {
	chanDemo()
}

/**
5,f
2,c
0,a
1,b
3,d
0,A
1,B
9,j
8,i
4,e
4,E
3,D
7,h
2,C
6,g
6,G
7,H
9,J
5,F
8,I

Process finished with the exit code 0


*/
