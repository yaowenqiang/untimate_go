package main
import "fmt"
func main() {
    friends := []string{
        "Annie",
        "Betty",
        "Charley",
        "Doug",
        "Edward",
    }

    fmt.Printf("%#v\n", friends)
    for _, v  := range friends {
        friends = friends[:2]
        fmt.Printf("v[%s]\n", v)
        fmt.Printf("%#v\n", friends)
    }

    fmt.Printf("%#v\n", friends)

    fmt.Println("*****************")

    friends2 := []string{
        "Annie",
        "Betty",
        "Charley",
        "Doug",
        "Edward",
    }

    fmt.Printf("%#v\n", friends2)
    for i := range friends2 {
        fmt.Printf("index [%d]\n", i)
        friends2 = friends2[:2]
        fmt.Printf("v[%s]\n", friends2[i])
        fmt.Printf("%#v\n", friends2)
    }
    fmt.Printf("%#v\n", friends2)
}
