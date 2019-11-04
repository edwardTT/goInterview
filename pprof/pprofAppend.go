// 展示内存增长和pprof，并不是泄露
package main

import (
    "fmt"
    "net/http"
    _ "net/http/pprof"
    "os"
    "time"
)

// 运行一段时间：fatal error: runtime: out of memory
func main() {
    // 开启pprof
    go func() {
        ip := "0.0.0.0:6060"
        if err := http.ListenAndServe(ip, nil); err != nil {
            fmt.Printf("start pprof failed on %s\n", ip)
            os.Exit(1)
        }
    }()


    //不要直接运行这段代码，他会引起内存大量泄漏。从而引起机器卡死
    //tick := time.Tick(time.Second / 100)
    //var buf []byte
    //for range tick {
    //    buf = append(buf, make([]byte, 1024*1024)...)
    //}
    var buf []byte
    for i := 1; i < 10; i++{
        buf = append(buf, make([]byte, 100*10)...)
        time.Sleep(1*time.Second)
    }
    fmt.Println("Quit the for loop")
    time.Sleep(1000*time.Second)
}
