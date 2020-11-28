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

type System struct {
    Xenia
    Pillar
}

func pull(x *Xenia, data []Data) (int, error) {
    for i := range data {
        if err := x.Pull(&data[i]); err != nil {
            return i, err
        }
    }
    return len(data), nil
}

func store(p *Pillar, data []Data) (int, error) {
    for i := range data {
        if err := p.Store(&data[i]); err != nil {
            return i, err
        }
    }

    return len(data), nil
}

func Copy(sys *System, batch int) error {
    data := make([]Data, batch)
    for {
        i , err := pull(&sys.Xenia, data)
        if i > 0 {
            if _ , err := store(&sys.Pillar, data); err != nil {
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
    sys := System{
        Xenia: Xenia{
            Host: "localhost:8000",
            Timeout: time.Second,
        },
        Pillar: Pillar{
            Host: "localhost:9000",
            Timeout: time.Second,
        },
    }

    if err := Copy(&sys, 3); err != io.EOF {
        fmt.Println(err)
    }
}
