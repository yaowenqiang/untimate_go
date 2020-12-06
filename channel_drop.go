package main
import (
    "fmt"
)

func drop() {
    const cap = 100
    ch := make(chan string, cap)

    go func() {
        for p := range ch {
            fmt.Println("employee : recv'd signal: ", p)
        }
    }()

    const work = 2000

    for w := 0; w < work; w++ {
        select {
        case ch <- "paper":
            fmt.Println("manager: send signal", w)
        default:
            fmt.Println("manager: dropped data", w)
        }
    }

    close(ch)
    fmt.Println("manager : send shutdown signal!")
}

func main() {
    drop()
}

