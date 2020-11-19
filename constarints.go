package main
import ("fmt")
const ui = 12345
const uf = 3.1415926
const ti int = 12345
const tf float64 = 3.141592

//const myUnit8 uint8 = 1000

var answer = 3 * 0.333

const third = 1 / 3.0

const zero = 1 / 3

const one int8 = 1

const two = 2 * one

const (
    maxInt = 9223372036854775807

    bigger = 9223372036854775808543522345

    //biggerInt int64 = 9223372036854775808543522345
)

func main () {
    const (
        A1 = iota 
        B1 = iota 
        C1 = iota 
    )
    fmt.Println("1:", A1, B1, C1)

    const (
        A2 = iota 
        B2
        C2
    )
    fmt.Println("2:", A2, B2, C2)

    const (
        A3 = iota + 1
        B3
        C3
    )
    fmt.Println("3:", A3, B3, C3)

    const (
        Ldate = 1 << iota
        Ltime
        Lmicroseconds
        Llongfile
        Lshortfile
        LUTC
    )

    fmt.Println("Log:",Ldate, Ltime, Lmicroseconds, Llongfile, Lshortfile, LUTC )
}

