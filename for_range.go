package main
import "fmt"

func main() {
    friends := [5]string{"Annie", "Betty", "Charley","Doug", "Edward"}
    fmt.Printf("Bfr[%s] : ", friends[1])

    for i := range friends {
        friends[i] = "Jack"
        if i == 1 {
            fmt.Printf("Aft[%s]\n", friends[1])
        }
    }

    friends = [5]string{"Annie", "Betty", "Charley","Doug", "Edward"}
    fmt.Printf("Bfr[%s] : ", friends[1])
    for i, v := range friends {
        friends[i] = "Jack"
        if i == 1 {
            fmt.Printf("v[%s]\n", v)
        }
    }

}
