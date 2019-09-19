package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	a uint16
	b uint32
	//	c uint16
}

type nilS struct{}

type nilSTest struct{}

func (s *nilS) addr() { fmt.Printf("nilS address %p\n", s) }

func (s *nilSTest) addr() { fmt.Printf("nilSTest address %p\n", s) }

type NILSB struct {
	A struct{}
	B struct{}
}

func main() {
	var s S
	var aUint16 uint16
	var bUint32 uint32
	var sString string
	var nils1 nilS
	var nils2 nilS
	var nilsb NILSB
	var nilSTest1 nilSTest
	var x [1000]struct{}
	var nilStructSlice = make([]struct{}, 100)
	var y = nilStructSlice[:50]
	anilStruct1 := struct{}{}
	bnilStruct2 := struct{}{}

	fmt.Printf("s unsafe size %d \n", unsafe.Sizeof(s))
	fmt.Printf("aUint16 unsafe size %d \n", unsafe.Sizeof(aUint16))
	fmt.Printf("bUint32 unsafe size %d \n", unsafe.Sizeof(bUint32))
	fmt.Printf("sString unsafe size %d \n", unsafe.Sizeof(sString))
	fmt.Printf("nils1 unsafe size %d \n", unsafe.Sizeof(nils1))
	nils1.addr()
	nils2.addr()
	nilSTest1.addr()
	fmt.Printf("nilsb unsafe size %d \n", unsafe.Sizeof(nilsb))
	fmt.Printf("nil struct array unsafe size %d \n", unsafe.Sizeof(x))
	fmt.Printf("nilStructSlice len y %d captation y %d \n", len(y), cap(y))
	fmt.Println(anilStruct1 == bnilStruct2)

	// $ go run dataSize.go
	// s unsafe size 8
	// aUint16 unsafe size 2
	// bUint32 unsafe size 4
	// sString unsafe size 16
	// nils1 unsafe size 0
	// nilS address 0x5851b0
	// nilS address 0x5851b0
	// nilSTest address 0x5851b0
	// nilsb unsafe size 0
	// nil struct array unsafe size 0
	// nilStructSlice len y 50 captation y 100
	// true
	
	
}
