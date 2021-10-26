package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter, request *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func errNoPermission(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func errUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown error")
}

func noError(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h       appHandler
		code    int
		message string
	}{
		{errPanic, 500, "Internal Server Error"},
		{errUserError, 400, "user error"},
		{errNotFound, 404, "Not Found"},
		{errNoPermission, 403, "Forbidden"},
		{errUnknown, 500, "Internal Server Error"},
		{noError, 200, "no error"},
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
			t.Errorf("expect (%d, %s)"+
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

API server listening at: 127.0.0.1:50219
=== RUN   TestErrWrapper
--- PASS: TestErrWrapper (0.00s)
PASS

Debugger finished with the exit code 0

API server listening at: 127.0.0.1:50243
=== RUN   TestErrWrapper
--- PASS: TestErrWrapper (0.00s)
PASS

Debugger finished with the exit code 0

*/
