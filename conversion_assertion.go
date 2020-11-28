// https://golang.org/src/io/io.go?s=12980:13040#L357}

package main
import (
"fmt"
"math/rand"
"time"
)


type Mover interface {
    Move()
}
type Locker interface {
    Lock()
    Unlock()
}
type MoveLocker interface {
    Mover
    Locker
}

type bike struct{}

func (bike) Move() {
    fmt.Println("Moving the boke")
}

func (bike) Lock() {
    fmt.Println("Locking the boke")
}


func (bike) Unlock() {
    fmt.Println("Unlocking the boke")
}

type car struct {}

func (car) String() string {
    return "Vroom!"
}


type cloud struct{}

func(cloud) String() string {
    return "Big data!"
}

type user struct {
    name string
    email string
}

//func (u user) String() string {
func (u *user) String() string {
    return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {
    var ml MoveLocker
    //var m Mover


    ml = bike{}
    fmt.Println("%#v", ml)
    //ml = m

    //b := m.(bike)
    //b2 , ok := m.(bike)

    //m2 := b
    //m3 := b2
    //fmt.Println("%#v", m2)


    mvs := []fmt.Stringer{
        car{},
        cloud{},
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 10; i++ {
        rn := rand.Intn(2)

        if v, is := mvs[rn].(cloud); is {
            fmt.Println("God Lucky:", v)
            continue
        }

        fmt.Println("Got Unlucky")
    }


    u := user{
        name: "jack",
        email: "jack@jack.com",
    }

    fmt.Println(u)
    fmt.Println(&u)
}




