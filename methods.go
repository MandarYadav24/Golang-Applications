package main
import "fmt"

type rect struct{
    length int
    width int
}

func (r *rect)area() int{
    return r.length * r.width
}

func (r rect)perimeter() int {
    return 2*r.length + 2*r.width
}

func main() {
  r := rect{length:30, width:20}
  fmt.Println("Area of rectangle:",r.area())
  fmt.Println("Perimeter of rectangle:",r.perimeter())
  
  rp := &r
  fmt.Println("Area of rectangle:",rp.area())
  fmt.Println("Perimeter of rectangle:",rp.perimeter())
}