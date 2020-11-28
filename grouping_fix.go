package main
import "fmt"

type Speaker interface {
    Speak()
}


type Dog struct {
    Name string
    IsMammal bool
    PackFactor int
}

func (d *Dog) Speak() {
    fmt.Printf(
        "Woof! My name is %s, it is %t I am a mammal with %d \n",
        d.Name,
        d.IsMammal,
        d.PackFactor,
    )

}

type Cat struct {
    Name string
    IsMammal bool
    ClimbFactor int
}


func (c *Cat) Speak() {
    fmt.Printf(
        "Meow! My name is %s, it is %t I am a mammal with %d \n",
        c.Name,
        c.IsMammal,
        c.ClimbFactor,
    )

}

func main() {
    animals := []Speaker{
        &Dog{
            Name: "Fido",
            IsMammal: true,
            PackFactor: 5,
        },
        &Cat{
            Name: "Tom",
            IsMammal: true,
            ClimbFactor: 10,

        },
    }

    for _, animal := range animals {
        animal.Speak()
    }

}
