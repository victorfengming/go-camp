package main

import (
	"fmt"
	"sync"
	//"time"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("@%d---%d\n", id, n)
		// 通知外面 做完了( channel 是一等公民)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// 告诉外面用的人 , 我这个channel怎么用
func createWorker(id int, wg *sync.WaitGroup) worker { // 告诉外面用的人 , 我这个channel怎么用
	//
	w := worker{
		in: make(chan int),
		// 函数式编程
		// 匿名函数来赋值
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
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
@1---98
@2---99
@2---67
@3---100
@3---68
@4---101
@4---69
@5---102
@5---70
@6---103
@6---71
@7---104
@7---72
@8---105
@8---73
@9---106
@9---74
@0---65
@1---66

Process finished with the exit code 0


*/
