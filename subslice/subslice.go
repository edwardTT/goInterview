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
		fmt.Printf("v value %d \n", v)
	}

	for i, value := range s {
		s[i] += 10
		fmt.Printf("Value: %d  Value_addr %X ElemAddr s %X \n", value, &value, &s[i])
	}
	for i, value := range ss {
		ss[i] += 10
		fmt.Printf("Value: %d  Value_addr %X ElemAddr ss__ %X \n", value, &value, &ss[i])
	}
	fmt.Println(s, ss)
	//Dump info
	//	v value 12
	//	v value 13
	//	v value 14
	//Value: 1  Value_addr C000022110 ElemAddr s C000018320
	//Value: 2  Value_addr C000022110 ElemAddr s C000018328
	//Value: 3  Value_addr C000022110 ElemAddr s C000018330
	//Value: 2  Value_addr C000022130 ElemAddr ss__ C000018340
	//Value: 3  Value_addr C000022130 ElemAddr ss__ C000018348
	//Value: 4  Value_addr C000022130 ElemAddr ss__ C000018350
	//	[11 12 13] [12 13 14]

}
