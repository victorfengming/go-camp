package main

import (
	"fmt"
	"time"
)

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("@%d---%d\n", id, n)
		// 通知外面 做完了( channel 是一等公民)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

// 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int) worker { // 告诉外面用的人 , 我这个channel怎么用
	//
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, w := range workers {
		w.in <- 'a' + i
		//res := <-workers[i].done
		//print(res)
	}
	for i, w := range workers {
		w.in <- 'A' + i
		//res := <-workers[i].done
		//print(res)
	}

	// wait for all of them
	for _, w := range workers {
		res := <-w.done
		res2 := <-w.done
		print(res)
		print(res2)
	}
	time.Sleep(time.Second)

}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()
}

/**
Channel as first-class citizen
@0---97
@4---101
@2---99
@3---100
@9---106
@5---102
@6---103
@8---105
@7---104
@1---98
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.chanDemo()
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/done/channel.go:45 +0x15d
main.main()
	E:/Projects/GolandProjects/go-camp/mooc/code/learngo/channel/done/channel.go:63 +0x57

goroutine 6 [chan send]:



*/
