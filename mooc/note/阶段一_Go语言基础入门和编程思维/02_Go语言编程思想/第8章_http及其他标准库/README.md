# 第8章_http及其他标准库

## 8-1 http标准库

![1635335363928](README/1635335363928.png)



 ```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response)

}

 ```



> 我现在想对 httpclinet进行控制
>
> 比如我想要访问手机版的imooc





```go
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {

	request, err := http.NewRequest(
		http.MethodGet,
		"http://www.imooc.com",
		nil,
	)
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", response)

}

```



![1635336316220](README/1635336316220.png)

![1635336328401](README/1635336328401.png)

## 8-2 json数据格式的处理



### +v格式化打印

```go
package main

import "fmt"

type Order struct {
	Id string
	Name string
	Quantity int
	TotalPrice int
}

func main() {
	o:=Order{
		Id:         "124",
		Name:       "lreag go",
		Quantity:   3,
		TotalPrice: 30,
	}

	fmt.Printf("%+v",o)

}

```





### 使用 json库 格式化库



```go
package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Id string
	Name string
	Quantity int
	TotalPrice int
}

func main() {
	o:=Order{
		Id:         "124",
		Name:       "lreag go",
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b,err:=json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",b)
}

```

### json字段

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	TotalPrice int`json:"total_price"`
}

func main() {
	o:=Order{
		Id:         "124",
		Name:       "lreag go",
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b,err:=json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n",b)
}
 
```

> 在结构体里面,首字母小写的字段是不能被看到的





### 省略空字段



```go
package main

import (
	"encoding/json"
	"fmt"
)

type Order struct {
	Id         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
}

func main() {
	o := Order{
		Id:         "124",
		//Name:       "lreag go",
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}
/**
{"id":"124","quantity":3,"total_price":30}

Process finished with the exit code 0
 */
```





### json 嵌套

```go
package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	Id         string `json:"id"`
	Item       OrderItem `json:"item"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
}

func main() {
	o := Order{
		Id: "124",
		Item: OrderItem{
			ID:    "8675645",
			Name:  "nancy",
			Price: 0,
		},
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

/**
{
  "id": "124",
  "item": {
    "id": "8675645",
    "name": "nancy",
    "price": 0
  },
  "quantity": 3,
  "total_price": 30
}

*/

```





### 指针也可以

```go
package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	Id         string `json:"id"`
	Item       *OrderItem `json:"item"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
}

func main() {
	o := Order{
		Id: "124",
		Item: &OrderItem{
			ID:    "8675645",
			Name:  "nancy",
			Price: 0,
		},
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

/**
{
  "id": "124",
  "item": {
    "id": "8675645",
    "name": "nancy",
    "price": 0
  },
  "quantity": 3,
  "total_price": 30
}

*/

```

### 切片也支持



```go
package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	Id         string `json:"id"`
	Item       *[]OrderItem `json:"item"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
}

func main() {
	o := Order{
		Id: "124",
		Item: &[]OrderItem{
			{
				ID:    "8675645",
				Name:  "nancy",
				Price: 50,
			},
			{
				ID:    "2456733",
				Name:  "yael",
				Price: 20,
			},
		},
		Quantity:   3,
		TotalPrice: 30,
	}

	//fmt.Printf("%+v",o)

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

/**
{
  "id": "124",
  "item": [
    {
      "id": "8675645",
      "name": "nancy",
      "price": 50
    },
    {
      "id": "2456733",
      "name": "yael",
      "price": 20
    }
  ],
  "quantity": 3,
  "total_price": 30
}

*/

```



![1635345615051](README/1635345615051.png)

## 8-3 第三方API数据格式的解析技巧



### map类型收 json

```go
package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
type Order struct {
	Id         string       `json:"id"`
	Item       []*OrderItem `json:"item"`
	Quantity   int          `json:"quantity"`
	TotalPrice int          `json:"total_price"`
}

func marshal() {
	o := Order{
		Id: "124",
		Item: []*OrderItem{
			{
				ID:    "8675645",
				Name:  "nancy",
				Price: 50,
			},
			{
				ID:    "2456733",
				Name:  "yael",
				Price: 20,
			},
		},
		Quantity:   3,
		TotalPrice: 30,
	}
	fmt.Println(o)
}
func unmarshal() {
	s := `{"id":"124","item":[{"id":"8675645","name":"nancy","price":50},{"id":"2456733","name":"yael","price":20}],"quantity":3,"total_price":30}
`
	var o Order
	json.Unmarshal([]byte(s), &o)
	fmt.Printf("%+v", o)
}


func parseNLP() {
	res := `{
"data": [
    {
        "synonym":"",
        "weight":"0.6",
        "word": "真丝",
        "tag":"材质"
    },
    {
        "synonym":"",
        "weight":"0.8",
        "word": "韩都衣舍",
        "tag":"品牌"
    },
    {
        "synonym":"连身裙;联衣裙",
        "weight":"1.0",
        "word": "连衣裙",
        "tag":"品类"
    }
]
}`
	//m := struct {
	//	Data []struct {
	//		Synonym string `json:"synonym"`
	//		Tag     string `json:"tag"`
	//	} `json:"data"`
	//}{}

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("-----------------------------------------------------------\n")
	fmt.Printf("%+v\n",m["data"].([]interface{})[2].(map[string]interface{})["synonym"])
	fmt.Printf("-----------------------------------------------------------\n")


	for k,v := range m{
		fmt.Println(k,v)
		for _,v2 := range v.([]interface{}){
			for v3,v4:=range v2.(map[string]interface{}){
				fmt.Printf("v3:%v---v4:%v\n",v3,v4)
			}
			//fmt.Printf("v1:%v---v2%v\n",v1,v2)

		}
	}


	fmt.Printf("%+v\n", m)
	//fmt.Printf("%+v, %+v\n", m.Data[2].Synonym, m.Data[2].Tag)
}


func main() {

	parseNLP()
}

```

其中的map[string]interface 属实复杂

```go
fmt.Printf("%+v\n",
		m["data"].([]interface{})[2].(map[string]interface{})["synonym"],
	)
```



### 结构体收json

> 换成结构体的

```go
m := struct {
		Data []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
		} `json:"data"`
	}{}
```

```go
err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v, %+v\n", m.Data[2].Synonym, m.Data[2].Tag)
```



## 8-4 gin框架介绍




## 8-5 为gin增加middleware



- gin-

### code 01 init

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

```



### code 02 加中间件拦截

```go
package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// middleware
	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		// 不管访问什么,都能先进到这里面来
		// log latency, response code
		logger.Info("incoming request:",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
		//log.Fatalf(c.Request.URL.Path)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello gin")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

```



### code03 requestId生成



```go
package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

func main() {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	// middleware
	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		// 不管访问什么,都能先进到这里面来
		// log latency, response code
		logger.Info("incoming request:",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
		//log.Fatalf(c.Request.URL.Path)
	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})
	r.GET("/ping", func(c *gin.Context) {

		h := gin.H{
			"message": "pong",
		}

		if rid, exists := c.Get("requestId"); exists {
			h["requestId"] = rid
		}
		c.JSON(200, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello gin")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

```





- gin-gonic/gin
- middleware的使用
- context的使用




## TODO For Blog



写一个脚本实现gitee page 自动更新

github page 自动同步

