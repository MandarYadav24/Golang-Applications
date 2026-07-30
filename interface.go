package main
import (
    "fmt"
    "math"
    )

type geometry interface{
    area() float64
    perimeter()float64
}

type rectangle struct{
    length, width float64
}

type circle struct{
    radius float64
}

func (r rectangle)area() float64{
    return r.length * r.width
}

func (r rectangle)perimeter() float64{
    return 2*r.length + 2*r.width
}

func (c circle)area() float64{
    return math.Pi * c.radius * c.radius
}

func (c circle)perimeter() float64{
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println("Area:",g.area())
    fmt.Println("Perimeter:",g.perimeter())
}

func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("Circle with radius:",c.radius)
    }
}

func main() {
  r := rectangle{length:12, width:10}
  c := circle{radius:20}
  fmt.Println("--------Rectangle--------")
  measure(r)
  fmt.Println("---------Circle----------")
  measure(c)
  
  detectCircle(r)
  detectCircle(c)
}