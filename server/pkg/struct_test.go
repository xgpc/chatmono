package pkg

import (
	"fmt"
	"testing"
)

type node struct {
	Name string `json:"name"`
	P    *node  `json:"p"`
}

func TestDemo(t *testing.T) {
	n := node{
		Name: "xxx",
		P:    nil,
	}

	n.P = &n

	body := StructToMap(n)
	fmt.Println(body)

}
