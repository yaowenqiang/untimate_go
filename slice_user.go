package main
import "fmt"

type user struct {
    likes int
}

func main() {
    users := make([]user, 3)

    shareUser := &users[1]

    shareUser.likes++
    for i := range users {
        fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
    }

    users = append(users, user{})
    shareUser.likes++

    for i := range users {
        fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
    }
}
