package main
import ( "fmt" )

type example struct {
    flag bool
    counter int16
    pi float32
}

func main() {
    var a int
    var b string
    var c float64
    var d bool
    fmt.Printf("var a int \t %T [%v]\n", a, a)
    fmt.Printf("var b string \t %T [%v]\n", b, b)
    fmt.Printf("var c float64 \t %T [%v]\n", c, c)
    fmt.Printf("var d bool \t %T [%v]\n", d, d)

    aa := 10
    bb := "hello"
    cc := 3.14159
    dd := true

    fmt.Printf("var aa int \t %T [%v]\n", aa, aa)
    fmt.Printf("var bb string \t %T [%v]\n", bb, bb)
    fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc)
    fmt.Printf("var dd bool \t %T [%v]\n", dd, dd)

    aaa := int32(10)
    fmt.Printf("var aaa int32(10) \t %T [%v]\n", aaa, aaa)

    var e1 example

    fmt.Printf("%+v\n", e1)

    e2 := example {
        flag: true,
        counter: 10,
        pi: 3.141592,
    }

    fmt.Println("Flag", e2.flag)
    fmt.Println("Counter", e2.counter)
    fmt.Println("Pi", e2.pi)


    //Declare a variable of an anonymous type set to its zero value
    var e11 struct {
        flag bool
        counter int16
        pi float32
    }

    //Display the value
    fmt.Printf("%+v\n", e11)

    //Declare a variable of an anonymous type and init using a struct iteral


    e22 := struct {
        flag bool
        counter int16
        pi float32
    } {
        flag: true,
        counter: 10,
        pi: 3.1415926,
    }
    fmt.Printf("%+v\n", e11)
    fmt.Printf("%+v\n", e22)

    fmt.Println("hello world")
}
