package b

import (
    "fmt"
    _ "goInterview/initStep/a"
)

func init(){
    fmt.Println("init in package b")
}

func BDump(){
    fmt.Printf("BDump \n")
}