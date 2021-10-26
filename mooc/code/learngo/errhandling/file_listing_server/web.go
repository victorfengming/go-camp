package main

import (
	//"io/ioutil"
	"learngo/errhandling/file_listing_server/file_listing"
	"net/http"
	"os"
	//"os"
)

func main() {
	http.HandleFunc("/list/", errWrapper(file_listing.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK
			// 错误处理
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}
