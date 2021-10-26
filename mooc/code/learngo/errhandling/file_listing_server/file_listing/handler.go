package file_listing

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError2 string

func (e userError2) Error() string {
	return e.Message()
}

func (e userError2) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//Path截取 切片 [ i:]
	if strings.Index(request.URL.Path, prefix) != 0 {
		// 说明没有以list开头
		return userError2("Path must start with " + prefix)
	}

	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		// 处理
		//http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
