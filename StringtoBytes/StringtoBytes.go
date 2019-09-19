package main

import (
	"fmt"
)

func main() {
	s := "123"
	ps := &s
	b := []byte(*ps)
	pb := &b

	s += "4"
	*ps += "5"
	b[1] = '0'

	fmt.Printf("ps address %P\n", ps)
	println(*ps)
	fmt.Printf("pb address %P\n", pb)
	println(string(*pb))
}
