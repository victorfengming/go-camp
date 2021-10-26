package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func main() {
	var h HelloService = hello{
		endpoint: "http://localhost:8080/",
	}
	msg, err := h.SayHello("够浪")
	if err != nil {
		log.Fatalf("%s", err)
		return
	}
	fmt.Print(msg)
	PrintFuncName(h)
}

type HelloService interface {
	SayHello(name string) (string, error)
}

type hello struct {
	endpoint  string
	FuncField func()
}

func (h hello) SayHello(name string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(h.endpoint + name)
	if err != nil {
		log.Fatalf("%s", err)
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
		return "", err
	}
	return string(data), nil

}

func (h hello) GetOrder(name string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(h.endpoint + name)
	if err != nil {
		log.Fatalf("%s", err)
		return "", err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
		return "", err
	}
	return string(data), nil

}

// 远程调用的本质
//

// val interface{} >>> java 的 Object对象
// 跟对象
func PrintFuncName(val interface{}) {
	t := reflect.TypeOf(val)
	t2 := reflect.ValueOf(val)
	num := t.NumMethod()
	for i := 0; i < num; i++ {
		f := t2.Field(i)
		if f.CanSet() {
			fmt.Println("aaa")
		}
		m := t.Method(i)
		fmt.Println(m.Name)
	}
	//t.MethodByName()
}
