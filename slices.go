package main

import (
    "fmt"
    "slices"
    )
    
func main() {
    var s []string
    fmt.Println("Slice:",s, "Len:",len(s), "Cap:",cap(s))
    
    s = make([]string, 3)
    fmt.Println("Slice:",s, "Len:",len(s), "Cap:",cap(s))
    
    s[0] = "Golang"
    s[1] = "Python"
    s[2] = "C"
    
    fmt.Println("Slice:",s, "Len:",len(s), "Cap:",cap(s))
    
    s = append(s, "C++")
    s = append(s, "Rust", "Swift","Ruby")
    fmt.Println("Slice:",s, "Len:",len(s), "Cap:",cap(s))
    
    c := make([]string, len(s))
    copy(c,s)
    fmt.Println("Slice:",c, "Len:",len(c), "Cap:",cap(c))
    
    l := s[2:5]
    fmt.Println("Slice:",l, "Len:",len(l), "Cap:",cap(l))
    
    k := s[:5]
    fmt.Println("Slice:",k, "Len:",len(k), "Cap:",cap(k))
    
    m := s[2:]
    fmt.Println("Slice:",m, "Len:",len(m), "Cap:",cap(m))
    
    n := []string{"A", "B", "C"}
    fmt.Println("Slice: ",n)
    
    q := []string{"D","E"}
    fmt.Println("Slice: ",q)
    
    if slices.Equal(n,q){
        fmt.Println("n==q")
    } else {
        fmt.Println("n!=q")
    }
    
    twoD := make([][]int, 3)
    for i := range 3{
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen{
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}