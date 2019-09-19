package main

import (
	"fmt"
	"sort"
)

type S struct {
	v int
}

func main() {
	s := []S{{1}, {3}, {5}, {2}}
	s2 := []S{{1}, {9}, {5}, {2}}
	sort.Slice(s, func(i, j int) bool { return s[i].v < s[j].v })
	fmt.Printf("%#v", s)
    println()
	sort.SliceStable(s2, func(i, j int) bool { return s2[i].v < s2[j].v })
	fmt.Printf("%#v", s2)
}
