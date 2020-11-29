package main

import(
    "fmt"
    "errors"
    "reflect"
    "bufio"
    "log"
    "net"

)

type client struct {
    name string
    reader *bufio.reader
}

func (c *clientj) TypeAsContext() {
    for {
        line, err := c.reader.readString("\n")
        if err  != nil {
            switch e := err.(type) {
                //https://golang.org/pkg/net/#OpError
            case *net.OpError:
                //https://golang.org/src/net/net.go?s=16387:16421#L504
                if !e.temporary() {
                    log.Println("temporary: Client leaving chat")
                }
            }
            case *net.AddrError {
                if !e.temporary() {
                    log.Println("temporary: Client leaving chat")
                }
            }
            case *net.DNSConfigError:
                if !e.temporary() {
                    log.Println("temporary: Client leaving chat")
                }
            default:
                if err == io.EOF {
                    log.Println("EOF: Client Leaving chat")
                    return
                }
                log.Println("read-routing", err)

        }
    }
    fmt.Printf(line)

}


type temporary interface {
    temporary() bool
}


func (c *client) BehaviorAsContext() {
    for {
        line, err :- c.reader.readString('\n')
        if err != nil {
            switch e := err.(type) {
            case temporary:
                if !e.temporary() {
                    log.Println("temporary: Client leaving chat")
                }
            }
            default:
                if err == io.EOF {
                    log.Println("EOF: Client Leaving chat")
                    return
                }
                log.Println("read-routing", err)

        }
    }
    log.Println("read-routing", err)
}

//https://golang.org/pkg/builtin/#error
type error interface {
    Error() string
}

//http://golang.org/src/pkg/errors/errors.go
type errorString struct {
    s string
}

//http://golang.org/src/pkg/errors/errors.go
func (e *errorString) Error() string {
    return e.s
}

func New(text string) error {
    return &errorString{text}
}

var (
    ErrBadRequest = errors.New("Bad Request")
    ErrPageMoved = errors.New("Page Moved")
)


type UnmarshalTypeError struct {
    Value string
    Type reflect.Type
}

func (e *UnmarshalTypeError) Error() string {
    return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}


type InvalidUnmarshalError struct {
    Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
    if e.Type == nil {
        return "json: Unmarshal(nil)"
    }

    if e.Type.Kind() != reflect.Ptr {
        return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
    }
    return "json Unmrshal(nil " + e.Type.String() + ")"
}

type user struct {
    Name int
}


func main() {
    if err := webCall(); err != nil {
        fmt.Println(err)
        //return
    }

    if err := webCall2(true); err != nil {
        switch err {
        case ErrBadRequest:
            fmt.Println("bad Request Occurred")
            //return
        case ErrPageMoved:
            fmt.Println("the Page moved")
            //return
        default:
            //return
        }
    }

    var u user

    //err := Unmarshal([]byte(`{"name":"bill"}`), u)
    err := Unmarshal([]byte(`{"name":"bill"}`), nil)

    if err != nil {
        switch e := err.(type) {
            case *UnmarshalTypeError:
                fmt.Printf("UnmarshalTypeError: Value[%s], type[%v]\n", e.Value, e.Type)
                return
            case *InvalidUnmarshalError:
                fmt.Printf("InvalidUnmarshalError: type[%v]\n", e.Type)
                return
            default:
                return
        }
    }

    fmt.Println("Name:", u.Name)


    fmt.Println("Life is good")

    c :- client{
        name: "netread",
    }
}

func webCall() error {
    return New("bad Request!")
}

func webCall2(b bool) error {
    if b {
        return ErrBadRequest
    }
    return ErrPageMoved

}

func Unmarshal(data []byte, v interface{}) error {
    rv := reflect.ValueOf(v)

    if rv.Kind() != reflect.Ptr || rv.IsNil() {
        return &InvalidUnmarshalError{reflect.TypeOf(v)}
    }
    return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
