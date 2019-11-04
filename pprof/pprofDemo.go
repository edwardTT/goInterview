package main

import (
    "log"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func main() {
    //// 开启pprof，监听请求
    //ip := "0.0.0.0:6060"
    //if err := http.ListenAndServe(ip, nil); err != nil {
    //    fmt.Printf("start pprof failed on %s\n", ip)
    //}
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    time.Sleep(100*time.Second)
}
