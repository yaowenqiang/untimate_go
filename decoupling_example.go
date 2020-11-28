package main
import (
    "fmt"
    "errors"
    "io"
    "math/rand"
    "time"
)

type Data struct {
    Line string
}


type Xenia struct {
    Host string
    Timeout time.Duration
}

type Pillar struct {
    Host string
    Timeout time.Duration
}


func (*Xenia) Pull(d *Data) error {
    switch rand.Intn(100) {
    case 1, 9:
        return io.EOF
    case 5:
        return errors.New("Error reading data from server")
    default:
        d.Line = "Data"
        fmt.Println("In:", d.Line)
        return nil
    }
}


func (*Pillar) Store(d *Data) error {
    fmt.Println("Out:", d.Line)
    return nil
}

/*
type System struct {
    Puller
    Storer
}
*/

type Puller interface {
    Pull(* Data) error
}

type Storer interface {
    Store(* Data) error
}

/*
type PullStorer interface {
    Puller
    Storer
}
*/

func pull(p Puller, data []Data) (int, error) {
    for i := range data {
        if err := p.Pull(&data[i]); err != nil {
            return i, err
        }
    }
    return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
    for i := range data {
        if err := s.Store(&data[i]); err != nil {
            return i, err
        }
    }

    return len(data), nil
}

func Copy(p Puller, s Storer, batch int) error {
    data := make([]Data, batch)
    for {
        i , err := pull(p, data)
        if i > 0 {
            if _ , err := store(s, data); err != nil {
                return err
            }
        }

        if err != nil {
            return err
        }
    }


}

func init() {
    rand.Seed(time.Now().UnixNano())
}
func main() {
    x := Xenia{
            Host: "localhost:8000",
            Timeout: time.Second,
    }
    p := Pillar{
        Host: "localhost:9000",
        Timeout: time.Second,
    }

    if err := Copy(&x, &p, 3); err != io.EOF {
        fmt.Println(err)
    }
}
