package main

import "fmt"

type user struct {
    name string
    email string
}

type notifier interface {
    notify()
}

/*
func (u *user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", 
            u.name,
            u.email)
}
*/

func (u user) notify() {
    fmt.Printf("Sending user email to %s<%s>\n", 
            u.name,
            u.email)
}

func (a administrator) notify() {
    fmt.Printf("Sending admin email to %s<%s>\n",
            a.name,
            a.email)
}

type admin struct {
    person user //NOT embedding
    level string
}

type administrator struct {
    user //embedding
    level string
}


func main() {
    ad := admin {
        person: user {
            name: "John smith",
            email: "jonh@yahoo.com",
        },
        level: "super",
    }

    ad.person.notify()


    ad2 := administrator {
        user: user {
            name: "administrator",
            email: "administrator@admin.com",
        },
        level: "top",
    }

    ad2.user.notify()
    //Inner type method is promoted
    ad2.notify()

    sendNotification(ad2)
    sendNotification(&ad2)
}

func sendNotification(n notifier) {
    n.notify()
}
