package main
import (
    "fmt"
    "sync"
    //"sync/atomic"
)

// go build -race
// go test -race
// go test -race -cpu 

var counter int32

func main() {
    const grs = 2
    var wg sync.WaitGroup

    wg.Add(grs)

    var mu sync.Mutex
    //var mu sync.RwMutex

    for i := 0; i < grs; i++ {
        go func() {
            for count := 0; count < 2; count++ {
                /*
                value := counter
                value++
                fmt.Println("logging")
                counter = value
                */
                //fmt.Println("logging")
                //atomic.AddInt32(&counter, 1)
                mu.Lock()
                {
                    value := counter
                    value++
                    //fmt.Println("logging")
                    counter = value
                }
                mu.Unlock()
            }
            wg.Done()
        }()
    }

    wg.Wait()

    fmt.Println("Final Counter:", counter)

}

