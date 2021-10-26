package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

func main() {
	var h = &hello{
		endpoint: "http://localhost:8080/",
	}
	msg, err := h.SayHello("够浪")
	if err != nil {
		log.Fatalf("%s", err)
		return
	}
	fmt.Print(msg)
	SetFuncField(h)
	_, _ = h.FuncField("nancy")

}

// val interface{} >>> java 的 Object对象
// 跟对象
func SetFuncField(val interface{}) {
	//t:= reflect.TypeOf(val)
	v := reflect.ValueOf(val) // zhizhen指针的反射
	ele := v.Elem()           // 指针指向的结构体
	t := ele.Type()           // 指针指向的结构体的类型信息
	num := t.NumField()       // 方法数量
	for i := 0; i < num; i++ {
		f := ele.Field(i)
		if f.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				name := args[0].Interface().(string)
				fmt.Printf("change method begin\n")
				client := http.Client{}
				resp, err := client.Get("http://localhost:8080/" + name)
				//resp, err := client.Get(addr)
				if err != nil {
					return []reflect.Value{
						reflect.ValueOf(""),
						reflect.Zero(reflect.TypeOf(err)),
					}
				}
				data, err := ioutil.ReadAll(resp.Body)

				if err != nil {
					return []reflect.Value{
						reflect.ValueOf(""),
						reflect.Zero(reflect.TypeOf(err)),
					}
				}
				fmt.Printf("change method end\n")
				return []reflect.Value{
					reflect.ValueOf(string(data)),
					reflect.Zero(reflect.TypeOf(new(error)).Elem()),
				}
			}
			f.Set(reflect.MakeFunc(f.Type(), fn))
		}
	}
}

func inn(addr string) func() ([]byte, error) {
	return func() ([]byte, error) {
		// 匿名函数
		fmt.Printf("change method begin\n")
		client := http.Client{}
		//resp, err := client.Get("http://localhost:8080/golang")
		resp, err := client.Get(addr)
		if err != nil {
			log.Fatalf("%s", err)
			return nil, err
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("%s", err)
			return nil, err
		}
		//print(data)
		//return string(data), nil
		fmt.Printf("change method end\n")
		return data, nil
	}
}

type HelloService interface {
	SayHello(name string) (string, error)
}

type hello struct {
	endpoint string
	// 只能改这个
	FuncField func(name string) (string, error)
	//getUser GetUser(req *UserReq)(*User,error)
}

func (h hello) SayHello(name string) (string, error) {
	return "", nil
}

func (h hello) GetOrder(name string) (string, error) {
	return "", nil
}
