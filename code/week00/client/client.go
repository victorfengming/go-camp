package main

import "net/http"

func main() {
	client := http.Client{}
	resp, err := client.Get("http://localhost:8080/victor")
	if err != nil {
		print(err)
	}
	print(resp.Status)
}
