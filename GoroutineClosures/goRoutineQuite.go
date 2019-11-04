package main

import (
    "fmt"
    "log"
    "math/rand"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    ok, quit := make(chan int, 1), make(chan int, 1)
    go func() {
        i := 0
        for {
            select {
            case <-quit:
                ok <- 1
                //return
            default:
                HeavyWork(i)
                i++
            }
        }
    }()
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    time.Sleep(5 * time.Second)
    quit <- 10
    <-ok

    time.Sleep(60*time.Second)
}

func HeavyWork(id int) {
    rand.Seed(int64(id))
    interval := time.Duration(rand.Intn(3)+1) * time.Second
    time.Sleep(interval)
    fmt.Printf("HeavyWork %-3d cost %v\n", id, interval)
}