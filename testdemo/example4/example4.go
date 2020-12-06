package main
import (
    "log"
    "encoding/json"
    "net/http"
)

func main() {
    http.HandleFunc("/sendjson", sendJSON)
    log.Println("listener: Startted: Listening on http://localhost:4000")
    http.ListenAndServe(":4000", nil)
}

func sendJSON(rw http.ResponseWriter, r *http.Request) {
    u := struct {
        Name string
        Email string
    }{
        Name: "Bill",
        Email: "bill@bill.com",
    }
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(200)
    json.NewEncoder(rw).Encode(&u)
}
