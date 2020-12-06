package main
import (
    "fmt"
    "time"
    "runtime"
    "math/rand"
)

func fanOutSem() {
    emps := 2000
    ch := make(chan string, emps)

    g := runtime.NumCPU()
    sem := make(chan bool, g)

    for e := 0; e <= emps; e++ {
        go func(emp int) {
            sem <- true
            {
                time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
                ch <- "paper"
                fmt.Println("manager : end signal", e)
            }
            <-sem
        }(e)
    }

    for emps >0 {
        p := <-ch
        emps--
        fmt.Println(p)
        fmt.Println("manager : recv'd signal: ", emps)
    }
}

func main() {
    fanOutSem()
}


