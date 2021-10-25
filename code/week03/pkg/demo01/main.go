package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("rsat")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, gopherCon SG")
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// 为了让他不死,就要写个死循环
	select {}
}
