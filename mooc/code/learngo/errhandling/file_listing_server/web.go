package main

import (
	log "github.com/sirupsen/logrus"
	"learngo/errhandling/file_listing_server/file_listing"
	//"io/ioutil"
	"net/http"
	"os"
	//"os"
)

func main() {
	http.HandleFunc("/", errWrapper(file_listing.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			log.Printf("Painic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request:%s"+
				"", err.Error())
			code := http.StatusOK
			// 错误处理
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}

	}
}
