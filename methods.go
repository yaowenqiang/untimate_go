package main
import "fmt"


type data struct {
    name string
    age int
}

func (d data) displayName() {
    fmt.Printf("%p My Nme is\n",d.name,  d.name)
}


func (d *data) setAge(age int) {
    d.age = age
    fmt.Println(d.name, "is Age", d.age)
}

func main() {
    d := data{
        name: "Bill",
    }

    fmt.Println("Proper calls to Methods:")

    d.displayName()
    d.setAge(50)

    //This is what Go is doing underneath
    //data.displayName(d)
    //(*data).setAge(&d, 45)

    f1 := d.displayName
    fmt.Printf("%p\n", f1)
    fmt.Printf("%p\n", d.displayName)
    f1()

    d.name = "jack"

    f1()


    f2 := d.setAge

    f2(45)

    d.name = "Jerry"

    f2(45)

}
