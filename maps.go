package main
import (
"fmt"
"sort"
)
type user struct {
    name string
    surname string
}
type users []user

type player struct {
    name string
    score int
}

func main() {
    //users := map[string]user{}
    users := make(map[string]user)
    users["Roy"] = user{"Rob", "Roy"}
    users["Ford"] = user{"Henry", "Ford"}
    users["Mouse"] = user{"Mickey", "Mouse"}
    users["Jackson"] = user{"Michael", "Jackson"}

    mouse := users["Mouse"]

    fmt.Printf("%+v\n", mouse)

    users["Mouse"] = user{"Jerry", "Mouse"}
    fmt.Printf("%+v\n", users["Mouse"])

    delete(users, "Roy")

    fmt.Println(len(users))

    delete(users, "Roy")
    fmt.Println("bye.")


    scores := make(map[string]int)

    score := scores["anna"]


    fmt.Println("Score:", score)

    score, ok := scores["anna"]

    fmt.Println("Score:", score, "Present:", ok)

    scores["anna"]++

    if n, ok := scores["anna"] ; ok {
        scores["anna"] = n + 1
    } else {
        scores["anna"] = 1
    }

    score, ok = scores["anna"]

    fmt.Println("Score:", score, "Present:", ok)

    //u := make(map[users]int)

   // for key, value := range u {
    //    fmt.Println(key, value)
   // }

    users2 := map[string]user {
        "Roy": {"Rob", "Roy"},
        "Ford": {"Henry", "Ford"},
        "Mouse": {"Mickey", "Mouse"},
        "Jackson": {"Michael", "Jackson"},
    }

    // map range are random
    for key ,value := range users2 {
        fmt.Println(key, value)
    }

    for key ,value := range users2 {
        fmt.Println(key, value)
    }

    fmt.Println("--------------Sort maps ----------")
    // sort map

    var keys []string

    for key := range users2 {
        keys = append(keys, key)
    }

    sort.Strings(keys)

    for _, key := range keys {
        fmt.Println(key, users2[key])
    }


    fmt.Println()

    players := map[string]player {
        "anan" : {"Anna", 42},
        "Jacob" : {"Jacob", 21},
    }

    //anna := &players['anna']
    //anna.score++

    player := players["anna"]
    player.score++
    players["anna"] = player

    scores2 := map[string]int {
        "anna" : 21,
        "Jacob" : 12,
    }


    double(scores2, "anna")

    fmt.Println("Score:", scores2["anna"])

}

func double(scores map[string]int, player string) {
    scores[player] = scores[player] * 2
}
