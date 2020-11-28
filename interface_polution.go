package main
import "fmt"

type Server interface {
    Start() error
    Stop() error
    Wait() error
}


type server struct {
    host string
}

func NewServer(host string) Server {
    return &server{host}
}

func (s *server) Start() error {
    fmt.Println("Server Started")
    return nil
}

func (s *server) Stop() error {
    fmt.Println("Server Stopped")
    return nil
}
func (s *server) Wait() error {
    fmt.Println("Server is waiting")
    return nil
    return nil
}

func main() {
    srv := NewServer("localhost")
    srv.Start()
    srv.Stop()
    srv.Wait()
}
