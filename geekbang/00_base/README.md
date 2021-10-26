![1635214530127](README/1635214530127.png)

![1635214534211](README/1635214534211.png)

![1635215175692](README/1635215175692.png)





而不是

```go
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


```

> HTTP 协议 而 RPC是一种理念

![1635215288948](README/1635215288948.png)



看来我们需要一个指针

![1635216252902](README/1635216252902.png)

与Java对比,没有构造函数



![1635216284128](README/1635216284128.png)

![1635216385897](README/1635216385897.png)

![1635216543771](README/1635216543771.png)



## 不加指针的

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) ChangeName(newName string) {
	u.Name = newName
}

func (u User) ChangeAge(newAge int) {
	u.Age = newAge

}

func main() {
	u:=User{
		Name: "TOm",
		Age:  18,
	}

	u.ChangeName("Jerry")
	u.ChangeAge(17)

	fmt.Println(u.Name)
	fmt.Println(u.Age)


}
```

```cmd
TOm
18

Process finished with the exit code 0
```



## 带上指针的

```go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) ChangeName(newName string) {
	u.Name = newName
}

func (u *User) ChangeAge(newAge int) {
	u.Age = newAge

}

func main() {
	u:=User{
		Name: "TOm",
		Age:  18,
	}

	u.ChangeName("Jerry")
	u.ChangeAge(17)

	fmt.Println(u.Name)
	fmt.Println(u.Age)


}
```

```cmd
Jerry
17

Process finished with the exit code 0
```

![1635217142338](README/1635217142338.png)

![1635217194184](README/1635217194184.png)

![1635217322890](README/1635217322890.png)

