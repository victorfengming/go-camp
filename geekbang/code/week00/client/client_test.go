package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetFuncField(t *testing.T) {
	fmt.Println("arst")

	helloService := &hello{
		endpoint:  "",
		FuncField: nil,
	}

	SetFuncField(helloService)

	res, err := helloService.SayHello("beeggo")
	if err != nil {
		t.FailNow()
	}
	if res != "hello golang" {
		t.FailNow()
	}

	assert.New(t)
	assert.Nil(t, hello{}, "srat")

}
