package file_listing

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	//Path截取 切片 [ i:]
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
