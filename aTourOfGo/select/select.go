//The select statement lets a goroutine wait on multiple communication operations.
//    A select blocks until one of its cases can run, then it executes that case. It
//chooses one at random if multiple are ready.
//select 语句使一个 Go 程可以等待多个通信操作。
//
//select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	var hartStep int = 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Printf("\n c case hartStep %d \n", hartStep)
		case <-quit:
			fmt.Println("Quit")
			fmt.Printf("\n quit case hartStep %d \n", hartStep)
			return
			//增加defaul时，hartstep是1000多。
			// 不用default时，hartStep是10。
			// 这也验证了，select语句在没有要接受的数据时会阻塞，不过有了default分支后就不会阻塞。
			//default:
			//	fmt.Println("Select default")
		}
		hartStep++
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	var global int
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-c)
			global++
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}
