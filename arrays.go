package main
import "fmt"

func main() {
    var fruits [5]string
    fruits[0] = "Apple"
    fruits[1] = "Orange"
    fruits[2] = "Banana"
    fruits[3] = "Grape"
    fruits[4] = "Plum"

    for i, fruit := range fruits {
        fmt.Println(i, fruit)
    }
}
