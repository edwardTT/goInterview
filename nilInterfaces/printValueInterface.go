package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

func (u *user) notify(){
	fmt.Printf("User %s email is %s \n", u.name,u.email)
}

func main(){
    u := user{"Edward Liu",
    	      "bwxzing@163.com"}

    sentNotifyInfo(&u)

}
func sentNotifyInfo(n notifier){
	n.notify()
}