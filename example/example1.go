package  main
import (
    "fmt"
    //GOPATH/src/github.com/ardanlabs/gotraining/topics/go/language/exporting/example1/counters
    //VENDOR/github.com/ardanlabs/gotraining/topics/go/language/exporting/example1/counters
    "counters"
)

func main() {
    counter := counters.AlertCounter(10)
    fmt.Println("Counter: %d\n", counter)
}
