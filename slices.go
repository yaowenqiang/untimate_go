package main
import "fmt"
func main() {
    fruits := make([]string, 5)
    fruits[0] = "Apple"
    fruits[1] = "Orange"
    fruits[2] = "Banana"
    fruits[3] = "Grape"
    fruits[4] = "Plum"

    //error you canj't access an index of a slice beyond its length
    //fruits[5] = "Runtime Error"

    fmt.Println(fruits)


    fruits2 := make([]string, 5, 8)
    fruits2[0] = "Apple"
    fruits2[1] = "Orange"
    fruits2[2] = "Banana"
    fruits2[3] = "Grape"
    fruits2[4] = "Plum"

    inspectSlice(fruits2)


    var data []string
    //var data []string{}

    //var es struct{}
    //es := struct{}{}

    lastCap := cap(data)

    for record := 1; record <= 1e5; record++ {
        value := fmt.Sprintf("Rec: %d", record)
        data = append(data, value)

        if lastCap != cap(data) {
            //fmt.Printf("Last cap: %d\n", lastCap)
            capChg := float64(cap(data) - lastCap) / float64(lastCap) * 100
            //fmt.Printf("capChg: %f\n", capChg)
            fmt.Printf("Addr[%p]  Index[%-5d]\tCap[%-6d - %2f]\n",data, record, cap(data), capChg)
        }
        lastCap = cap(data)

    }

    fruits3 := make([]string, 5, 8)
    fruits3[0] = "Apple"
    fruits3[1] = "Orange"
    fruits3[2] = "Banana"
    fruits3[3] = "Grape"
    fruits3[4] = "Plum"

    inspectSlice(fruits3)

    slice4 := fruits3[2:4]
    slice5 := fruits3[2:4:4]

    inspectSlice(slice4)

    slice4[0] = "Changed"
    slice4 = append(slice4, "Changed with side wfffect")

    inspectSlice(fruits3)
    inspectSlice(slice4)


    inspectSlice(slice5)
    slice5 = append(slice5, "Changed without side wfffect")
    inspectSlice(slice5)


}

func inspectSlice(slice []string) {
    fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))

    for i, s := range slice {
        fmt.Printf("[%d]  %p %s\n",
        i,
        &slice[i],
        s)
    }
}
