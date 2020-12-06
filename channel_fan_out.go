package main
import (
    "fmt"
    "time"
    "math/rand"
)
func fanout() {
    emps := 2000
    ch := make(chan string, emps)
    for e:= 0; e < emps; e++ {
        go func(emp int) {
            time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
            ch <- "paper"
            fmt.Println("employee: send signal: " ,emp)
        }(e)
    }

    for emps > 0 {
        p := <-ch
        emps--
        fmt.Println(p)
        fmt.Println("manager: recv's signal: ", emps)
    }

    time.Sleep(time.Second)
    fmt.Println("-------------------------")
}

func main() {
    fanout()
}

