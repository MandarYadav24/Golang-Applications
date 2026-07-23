package main

import (
    "fmt"
    )
    
func main() {
    var a[5] int
    fmt.Println("Array a: ",a)
    
    a[4] = 200
    fmt.Println("Array a: ",a)
    
    b := [5]int {2,43,45,66,78}
    fmt.Println("Array b: ",b)
    
    b[3] = 4
    fmt.Println("Array b: ",b)
    
    c := [...]int{2,4,56,8,10}
    fmt.Println("Array c: ",c)
    
    c = [...]int{10,2:20,40,10}
    fmt.Println("Array c: ",c)
    
    var twoD [2][3] int
    for i := range 2 {
        for j := range 3 {
            twoD[i][j] = i+j
        }
    }
    fmt.Println("Two D array: ",twoD)
    
    twoD = [2][3] int{
        {1,2,3},
        {4,5,6},
    }
    fmt.Println("Two D array: ",twoD)
}