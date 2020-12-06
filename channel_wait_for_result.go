
package main

import (
    "fmt"
    "time"
    "math/rand"

)

func waitForResult() {
    ch := make(chan  string)

    go func() {
        time.Sleep(time.Duration(rand.Intn(500)) + time.Millisecond)
        ch <- "paper"
        fmt.Println("employee: send signal")
    }()

    p := <-ch

    fmt.Println("manager: recv'd signal: ", p)

    time.Sleep(time.Second)
    fmt.Println("-------------------------")

}

func main() {
    waitForResult()
}

