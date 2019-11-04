package main

import "fmt"

func add1() func() int {
	var count int

	return func() int {
		count++
		return count
	}
}

func main() {
	f1 := add1()
	f1()
	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())
}
