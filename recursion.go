package main
import (
    "fmt"
    )

func factorial(no int) int {
    if no == 0 {
        return 1
    }
    return no * factorial(no - 1)
}

func main() {
  result := factorial(7)
  fmt.Println(result)
  
  var fib func(no int) int
  
  fib = func(no int) int {
      if no < 2 {
          return no
      }
      
      return fib(no - 1) + fib(no - 2)
  }
  
  result1 := fib(7)
  fmt.Println(result1)
}