package main
import "fmt"

func main() {
  arr := []int{1,2,3}
  s1 := arr[1:]
  fmt.Println(arr)
  fmt.Println(s1)

  s2 := make([]int, 3,4)
  for i := range 3{
    s2[i] = 10 * (i+1)
  }

  fmt.Println(s2)
  s2 = append(s2,40)
  fmt.Println(s2)
  fmt.Println("length:",len(s2))
  fmt.Println("capacity:",cap(s2))

  s2 = append(s2,50)
  fmt.Println(s2)
  fmt.Println("length:",len(s2))
  fmt.Println("capacity:",cap(s2))

  m1 := make(map[int]string)
  m1[1] = "Apple"
  fmt.Println(m1)

  m2 := map[int]string{
    1: "Cat",
    2: "Dog",
  }
  fmt.Println(m2)

  name, ok := m2[3]; if !ok {
    fmt.Println("No key")
  } else {
    fmt.Println(name)
  }
}
