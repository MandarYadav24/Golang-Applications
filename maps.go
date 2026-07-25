package main

import (
    "fmt"
    "maps"
    )
    
func main() {
    m := make(map[string]int)
    m["Key1"] = 10
    m["Key2"] = 20
    fmt.Println("keys:",m)
    
    val1 := m["Key1"]
    fmt.Println("Value1:",val1)
    
    val3 := m["Key3"]
    fmt.Println("Value3:",val3)
    
    fmt.Println("Length:",len(m))
    
    delete(m, "Key2")
    fmt.Println("Map after delete:",m,"Length:",len(m))
    
    clear(m)
    fmt.Println("Map after clear, To clear everything",m)
    
    _, v := m["Key2"]
    fmt.Println("Present:",v)
    
    m["Key2"] = 20
    _, v = m["Key2"]
    fmt.Println("Present:",v)
    
    m1 := map[string]int {"Apple":1,
    "Banana":2, "Orange":3}
    fmt.Println(m1)
    
    m2 := map[string]int {"Apple":1,
    "Banana":2, "Orange":3}
    
    if maps.Equal(m1,m2){
        fmt.Println("map1 is equal to map2")
    }
    
}