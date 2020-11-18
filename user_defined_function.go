package main
import ("fmt")

type user struct {
    name string
    email string
}

func main() {
    u1 := createUserV1();
    u2 := createUserV2();
    fmt.Println("u1", &u1, "u2", u2)
}

func createUserV1() user {
    u := user {
        name: "Bill",
        email: "bill@bill.com",
    }

    fmt.Println("V1", &u)
    return u
}

func createUserV2() *user {
    u := user {
        name: "Bill",
        email: "bill@bill.com",
    }

    fmt.Println("V2", &u)
    return &u
}
