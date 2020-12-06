package main
import (
    "fmt"
    "runtime"
    "sync"
    //"time"
    //"math/rand"
)

func fanOutBounded() {
    work := []string{ "papaer", "papaer", "papaer", "papaer", "papaer", 2000:"paper"}
    //fmt.Println(len(work))

    g := runtime.NumCPU()

    var wg sync.WaitGroup
    wg.Add(g)

    ch := make(chan string, g)

    for e := 0; e < g; e++ {
        go func(emp int) {
            defer wg.Done()
            for p := range ch {
                fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
            }
            fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
        }(e)
    }


    for _, wrk := range work {
        ch<- wrk
    }

    close(ch)
    wg.Wait()
}

func main() {
    fanOutBounded()
}

