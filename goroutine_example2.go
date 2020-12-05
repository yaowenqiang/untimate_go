package main
import (
    "crypto/sha1"
    "runtime"
    "strconv"
    "sync"
    "fmt"
)

// $ ./goroutine_example2 | cut -c1 | grep '[AB]' | uniq
// B: 49999: b01f85c46b2ff221fcb313f8e5215daa82600a3b
func init() {
    runtime.GOMAXPROCS(1)
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    fmt.Println("Create Goroutines!")

    go func() {
        printHashes("A")
        wg.Done()
    }()

    go func() {
        printHashes("B")
        wg.Done()
    }()

    fmt.Println("Waiting To Finish")
    wg.Wait()

    fmt.Println("Terminating Program!")

}

func printHashes(prefix string) {
    for i := 0; i <= 50000; i++ {
        num := strconv.Itoa(i)
        //fmt.Printf("%+v", num)
        sum := sha1.Sum([]byte(num))
        fmt.Printf("%s: %05d: %x\n", prefix, i , sum)
    }

    fmt.Println("Completed", prefix)
}
