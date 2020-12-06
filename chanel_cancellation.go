package main
import (
    "fmt"
    "context"
    "time"
    "math/rand"
)

func cancellation() {
    duration := 150 * time.Millisecond
    ctx, cancel := context.WithTimeout(context.Background(), duration)

    defer cancel()

    ch := make(chan string, 1)

    go func() {
        time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
        ch <-"paper"
    }()

    select {
    case d := <-ch:
        fmt.Println("work complete", d)
    case <-ctx.Done():
        fmt.Println("work canceled")
    }

    time.Sleep(time.Second)
    fmt.Println("---------------")
}

func main() {
    cancellation()
}

