package main

import (
	"fmt"
	"reflect"
)

type S struct {
	a, b, c string
}

func main() {
	x := interface{}(&S{"a", "b", "c"})
	y := interface{}(&S{"a", "b", "C"})
	fmt.Println(x == y)                  //A
	fmt.Println(reflect.DeepEqual(x, y)) //比较数据的类型和内容是否一致。
}
