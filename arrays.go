package main
import ("fmt")

const (
    rows = 2 * 1024
    cols = 2 * 1024
)

var matrix [rows][cols]byte

type data struct {
    v byte
    p *data
}

var list *data

func main() {
    var last *data
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            var d data
            if list == nil {
                list = &d
            }
            if last != nil {
                last.p = &d
            }

            last = &d

            if rows%2 == 0 {
                matrix[row][col] = 0xFF
                d.v = 0xFF
            }
        }

    }
}

func LinkedListTraverse() int {
    var ctr int
    d := list
    for d != nil {
        if d.v == 0xFF {
            ctr++
        }
        d = v.p
    }
    return ctr
}

func ColumnTraverse() int {
    var ctr int
    for col := 0; col < cols; col++ {
        for row := 0; row < rows; row++ {
            if matrix[row][col] == 0xFF {
                ctr++
            }
        }
    }
    return ctr
}

func RowTraverse() int {
    var ctr int
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if matrix[row][col] == 0xFF {
                ctr++
            }
        }
    }
    return ctr
}
