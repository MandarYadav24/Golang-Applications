package main

import (
    "fmt"
    )

func add(no1, no2 int) int {
    return no1 + no2
}

func doubleAdd(no1, no2, no3 int) int {
    return no1 + no2 + no3
}

func swap(no1, no2 int) (int, int) {
    return no2, no1
}

func sums(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, v := range nums{
        total = total + v
    }
    fmt.Println(total)
}
   
func main() {
    v1 := add(11,21)
    v2 := doubleAdd(11,21,51)
    
    fmt.Println("Add:",v1)
    fmt.Println("Double add:",v2)
    
    v3 := 20
    v4 := 40
    v3, v4 = swap(v3,v4)
    fmt.Println("value v3:",v3)
    fmt.Println("value v4:",v4)
    
    _,v5 := swap(100,200)
    fmt.Println("value v5:",v5)
    
    sums(11)
    sums(11,21)
    sums(11,21,51)
    
    nums := []int{11,21,51,101,121}
    sums(nums...)
}