package main
import "fmt"

type person struct {
    name string
    age int
}

func newPerson(name string) *person {
    p := person{name:name}
    p.age = 50
    return &p
}

func main() {
  p1 := newPerson("Pratik")
  fmt.Println(*p1)
  
  fmt.Println(person{"Jay",30})
  fmt.Println(person{name:"Vijay"})
  fmt.Println(person{age:40})
  
  p2 := person{name:"Aniket", age:20}
  fmt.Println(p2.age)
  fmt.Println(p2.name)
  
  p3 := &p2
  fmt.Println(*p3)
  
  p3.name = "Nikhil"
  fmt.Println(p3.name)
  
  car := struct {
      name string
      isParked bool
  }{
      "Fortuner",
      true,
  }
  fmt.Println(car)
}