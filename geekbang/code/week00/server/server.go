package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(HelloWorld(request.Method))
	fmt.Println("req> Path>", request.URL.Path)
	fmt.Fprintf(writer, "IT Works")
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListAndS:", err)
	}
}

func HelloWorld(name string) string {
	return name + "over"
}
