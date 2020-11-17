package main
import ("fmt")

func main() {
    count := 10
    fmt.Println("count:\tValue of[", count , "]\tAddr of[", &count, "]")
    increment(count)
    fmt.Println("count:\tValue of[", count , "]\tAddr of[", &count, "]")

    increment2(&count)
    fmt.Println("count:\tValue of[", count , "]\tAddr of[", &count, "]")
}

func increment(inc int) {
    inc++;
    fmt.Println("inc:\tValue of[", inc , "]\tAddr of[", &inc, "]")
}

func increment2(inc *int) {
    *inc++;
    fmt.Println("inc:\tValue of[", inc , "]\tAddr of[", &inc, "]\tValue Points to[", *inc, "]")
}
