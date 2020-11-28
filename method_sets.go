package main
import "fmt"

type notifier interface {
    notify()
}

type user struct {
    name string
    email string
}

type duration int

func (d *duration) notify() {
    fmt.Printf("Sedning Notification in ", *d)
}


func (u *user) notify() {
    fmt.Printf("Seding User Email TO %s<%s>\n",
        u.name,
        u.email)
}
func sendNotification(n notifier) {
    n.notify()
}

func main() {
    //duration(42).notify()
    u := user{"Bill", "bill@bill.com"}
    sendNotification(&u)
}

