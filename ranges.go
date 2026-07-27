package main
import (
    "fmt"
    )

func main() {
  nums := []int{12,11,224,345,56,78}
  sum := 0
  for i, num := range nums{
      fmt.Print(i," ")
      sum += num
  }
  fmt.Println("\n", sum)
  
  for i, num := range nums{
      if num == 345 {
          fmt.Println("Index of 345:", i)
          break
      }
  }
  
  m1 := map[int]string {
      1:"C", 2:"C++", 3:"Golang",
  }
  for k, v := range m1 {
      fmt.Printf("%d : %s \n", k,v)
  }
  
  for k := range m1 {
      fmt.Println("Key:",k)
  }
  
  for i, c := range "Golang" {
      fmt.Println(i,c)
  }
}