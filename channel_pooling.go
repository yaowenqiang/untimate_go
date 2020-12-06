package main
import (
    "fmt"
    "runtime"
)
func pooling() {
    ch := make(chan string)

    g := runtime.NumCPU()

    for e := 0; e < g; e++ {
        go func(emp int) {
            for p := range ch {
                fmt.Printf("employee: %d : recv'd signal : %s\n", emp, p)
            }

            fmt.Printf("employee: %d : recv shutdown signal\n", emp)
        }(e)
    }

    const work = 100

    for w := 0; w < work; w++ {
        ch <- "paper"
        fmt.Println("manager : end signal", w)
    }

    close(ch)

    fmt.Println("manager : send shutdown signal!")
}

func main() {
    pooling()
}

