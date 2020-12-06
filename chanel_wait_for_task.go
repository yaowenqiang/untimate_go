package main
import (
    "fmt"
    "time"
    "math/rand"
)

func waitForTask() {
    ch := make(chan string)
    go func() {
        p := <-ch
        fmt.Println("employee: recv's signal: ", p)
    }()

    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    ch <- "paper"
    fmt.Println("manager: send signal")
    time.Sleep(time.Second)
    fmt.Println("-----------------------------------------------")
}

func main() {
    waitForTask()
}

