package main

type S struct {
	name string
}

func main() {
	m := map[string]S{"x": S{"one"}}
	m["x"].name = "two"
}

//.\MapImmutability1.go:9:14: cannot assign to struct field m["x"].name in map
