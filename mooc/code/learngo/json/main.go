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

func main() {

	unmarshal()
}
