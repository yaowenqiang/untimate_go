package main
import "fmt"

type printer interface {
    print()
}

type user struct {
    name string
}

func (u user) print() {
    fmt.Printf("User Name: %s\n", u.name)
}

func main() {
    u := user{"Bill"}

    entities := []printer{
        u,
        &u,
    }
    fmt.Printf("%#v\n", entities)

    u.name =  "Bill_CHG"

    for _, e := range entities {
        e.print()
    }

    fmt.Println("Hello World")
    fmt.Println(12345)
    fmt.Println(3.14159)
    fmt.Println(true)

    myPrintln("Hello World")
    myPrintln(12345)
    myPrintln(3.14159)
    myPrintln(true)


}

func myPrintln (a interface{}) {
    //v, ok := a.(string)

    switch v := a.(type) {
    case string:
        fmt.Printf("Is string : type(%T) : value(%s)\n", v, v)
    case int:
        fmt.Printf("Is int : type(%T) : value(%d)\n", v, v)
    case float64:
        fmt.Printf("Is float64 : type(%T) : value(%f)\n", v, v)
    default:
        fmt.Printf("Is unknown : type(%T) : value(%v)\n", v, v)
    }
}
