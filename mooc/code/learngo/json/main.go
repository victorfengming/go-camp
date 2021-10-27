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
	fmt.Printf("%+v\n",
		m["data"].([]interface{})[2].(map[string]interface{})["synonym"],
	)
	fmt.Printf("-----------------------------------------------------------\n")

	for k, v := range m {
		fmt.Println(k, v)
		for _, v2 := range v.([]interface{}) {
			for v3, v4 := range v2.(map[string]interface{}) {
				fmt.Printf("v3:%v---v4:%v\n", v3, v4)
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
