package main
import "fmt"



type Server struct {
    host string
}

func NewServer(host string) *Server {
    return &Server{host}
}

func (s *Server) Start() error {
    fmt.Println("Server Started")
    return nil
}

func (s *Server) Stop() error {
    fmt.Println("Server Stopped")
    return nil
}
func (s *Server) Wait() error {
    fmt.Println("Server is waiting")
    return nil
}

func main() {
    srv := NewServer("localhost")
    srv.Start()
    srv.Stop()
    srv.Wait()
}
