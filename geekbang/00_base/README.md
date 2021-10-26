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

