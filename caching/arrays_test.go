package caching
import "testing"
var fa int

func BenchmarkLinkListTraverse(b *testing.B) {
    var a int
    for i := 0; i < b.N; i++ {
        a = LinkedListTraverse();
    }
    fa = a
}

func BenchmarkColumntraverse(b *testing.B) {
    var a int
    for i := 0; i < b.N; i ++ {
        a = ColumnTraverse()
    }
    fa = a
}

func BenchmarkRowtraverse(b *testing.B) {
    var a int
    for i := 0; i < b.N; i++  {
        a = ColumnTraverse()
    }
    fa = a
}
