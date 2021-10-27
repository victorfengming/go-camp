package main

import "fmt"

func main() {
	//fmt.Printf("hellow from goroutine %d")
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				fmt.Printf("hellow from goroutine %d\n", ii)
			}
		}(i)
	}
	select {}
}
