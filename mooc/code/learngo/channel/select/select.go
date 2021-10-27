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
	for n := range c {
		time.Sleep(1 * time.Second)
		fmt.Printf("@%d---%d\n", id, n)

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
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {

		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = work
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			// 如果超过 800 毫秒 之内没有生成数据
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("queue len is", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}

}

/**
queue len is 3
@0---0
queue len is 4
@0---0
queue len is 5
@0---1
queue len is 10
@0---1
queue len is 10
@0---2
queue len is 11
@0---2
queue len is 12
@0---3
queue len is 13
@0---3
queue len is 14
@0---4
queue len is 15
bye

Process finished with the exit code 0

*/
