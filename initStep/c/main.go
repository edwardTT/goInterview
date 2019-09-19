package main

import (
    "fmt"
    "goInterview/initStep/b"
)
func init(){
    fmt.Printf("Init main package \n")
}
func main(){
    fmt.Printf("run in main \n")
    b.BDump()
}