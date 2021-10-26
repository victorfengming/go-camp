package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

//func errPanic(writer http.ResponseWriter,request *http.Request) error{
//	panic(123)
//}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
	}

	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://localhost:8888/list/fib2.txt",
			nil,
		)
		f(response, request)
		b, _ := ioutil.ReadAll(response.Body)
		body := strings.Trim(string(b), "\n")
		if response.Code != tt.code ||
			body != tt.message {
			t.Errorf("expect (%d, %s);"+
				"expect (%d, %s);\"",
				tt.code,
				tt.message,
				response.Code,
				body,
			)
		}
	}
}

/**
=== RUN   TestErrWrapper
time="2021-10-26T20:55:25+08:00" level=info msg="Painic: 123"
--- PASS: TestErrWrapper (0.02s)
PASS

Process finished with the exit code 0
*/
