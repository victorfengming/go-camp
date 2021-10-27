package main

import (
	"fmt"
	"sync"
	//"time"
)

func doWork(
	id int, c chan int, wg *sync.WaitGroup,
) {
	for n := range c {
		fmt.Printf("@%d---%d\n", id, n)
		// 通知外面 做完了( channel 是一等公民)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

// 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int, wg *sync.WaitGroup) worker { // 告诉外面用的人 , 我这个channel怎么用
	//
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {

	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	wg.Add(20)

	for i, w := range workers {
		w.in <- 'a' + i
		//wg.Add(1)

	}
	for i, w := range workers {
		w.in <- 'A' + i
	}

	// wait for all of themtime.Sleep(time.Second)
	wg.Wait()

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
