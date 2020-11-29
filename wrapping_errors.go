package main
import (
    "fmt"
    "github.com/pkg/errors"
)

type AppError struct {
    State int
}

func (c *AppError) Error() string {
    return fmt.Sprintf("App Error, State: %d", c.State)
}

func main() {
    if err := firstCall(10); err != nil {
        switch v := errors.Cause(err).(type) {
        case *AppError:
            fmt.Printf("Custom App Error:", v.State)
        default:
            fmt.Println("Default Error")
        }

        fmt.Println("\ntack Trace\n*******************************")
        fmt.Printf("%+v\n", err)
        fmt.Printf("No Trace\n**********************************\n")
        fmt.Printf("%v\n", err)
    }
}

func firstCall(i int) error {
    if err := secondCall(i); err != nil {
        return errors.Wrapf(err, "firstCall->secondCall(%d)", i)
    }
    return nil
}

func secondCall(i int) error {
    if err := thirdCall(); err != nil {
        return errors.Wrap(err, "SecondCall->ThirdCall()")
    }
    return nil
}

func thirdCall() error {
    return &AppError{99}
}
