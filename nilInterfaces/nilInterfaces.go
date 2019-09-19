package main

import (
	"unsafe"
)

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func main() {
	//cannot convert nil to type S
	//println(InitType() == nil)
	var sNil struct{}
	println(unsafe.Sizeof(sNil)) // print 0
	println(InitPointer() == nil)
	println(InitEfaceType() == nil)
	println(InitEfacePointer() == nil)
	println(InitIfaceType() == nil)
	println(InitIfacePointer() == nil)
}
