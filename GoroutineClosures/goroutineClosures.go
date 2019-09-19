package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(j int) {
			defer wg.Done()
			mu.Lock()
			m[j] = j
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	println()
        for v, ok:= range m{
                    println(v)
                    println()
                    println(ok)
        }
}