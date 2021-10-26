package main

import (
	"errors"
	"fmt"
)

// Uixianui先实现一个方法,可以支持在切片的中间加入值

func main() {
	values := []interface{}{0, 1, 2, 3, 5}
	newV, err := Add(values, 4, 4)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	// want to print : {0,1,2,3,4,5}
	for i, value := range newV {
		fmt.Printf("下表:%d, 值:%d\n", i, value)
	}
}

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {

	if index < 0 || index > len(values) {
		return nil, errors.New("index is invalid!!!")
	}

	res := []interface{}{}

	// 放好 0.1.2.3
	//把4放进去
	// 把剩下的放进去
	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}
	res = append(res, val)
	for i := index; i < len(values); i++ {
		v := values[i]
		res = append(res, v)
	}

	return res, nil
}

func Delete(values []interface{}, index int) []interface{} {
	if index < 0 || index > len(values) {
		return values
	}

	res := []interface{}{}

	for i := index; i < len(values); i++ {
		if index == i {
			continue
		}
		v := values[i]
		res = append(res, v)
	}

	return res

}
