package main

import (
    "fmt"
)

type notifier interface {
    notify()
}

type user struct{
    name string
    email string
}

type admin struct {
    name string
    email string
    permission int
}

func (u *user) notify(){
    fmt.Printf("Dump the name %s \n", u.name)
}

func (u *admin) notify(){
    fmt.Printf("Dump the name %s \n", u.name)
}

func main (){
    var edward = user{
        name: "Edward Liu",
        email: "bwxzing@163.com",
    }

    var Evila = admin{
        name: "Evila",
        email: "evila@163.com",
        permission: 4,
    }

    SentNotify(&edward)

    SentNotify(&Evila)

}

func SentNotify(n notifier){
    n.notify()
}