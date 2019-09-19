package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	ss := s[1:]
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i, value := range s {
		s[i] += 10
		fmt.Printf("Value: %d  Value_addr %X ElemAddr s %X \n", value, &value, &s[i])
	}
	for i, value := range ss {
		ss[i] += 10
		fmt.Printf("Value: %d  Value_addr %X ElemAddr ss %X \n", value, &value, &s[i])
	}
	fmt.Println(s, ss)
	//Dump info
	// $ go run subslice.go
	// Value: 1  Value_addr C000086050 ElemAddr s C00009C020
	// Value: 2  Value_addr C000086050 ElemAddr s C00009C028
	// Value: 3  Value_addr C000086050 ElemAddr s C00009C030
	// Value: 2  Value_addr C000086088 ElemAddr ss C00009C020
	// Value: 3  Value_addr C000086088 ElemAddr ss C00009C028
	// Value: 4  Value_addr C000086088 ElemAddr ss C00009C030
	// [11 12 13] [12 13 14]

}
