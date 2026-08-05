package main
import "fmt"

func SliceIndex[S ~[]E, E comparable](s S, v E) int{
    for i := range s {
        if v == s[i]{
            return i
        }
    }
    return -1
}

type List[T any] struct{
    head, tail *element[T]
}

type element[T any] struct{
    next *element[T]
    val T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func main() {
  var s = []string {"Golang","kafka","redis"}
  fmt.Println("Index of kafka:",SliceIndex(s,"kafka"))
  
  var i = []int {12,14,18,20}
  fmt.Println("Index of 18:",SliceIndex(i,18))
  
  _ = SliceIndex[[]string, string](s,"redis")
  
  lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)
    fmt.Println("list:", lst.AllElements())
}