package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Printf("hellow from goroutine %d")
	var a [1000]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			// race condition
			for {
				a[i]++
				//fmt.Printf("hellow from goroutine %d\n", ii)
				// 交出控制权
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}

/**


 */
