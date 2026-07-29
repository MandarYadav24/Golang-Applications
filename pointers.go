package main
import "fmt"

func noChangeVal(no int) {
    no = 0
}

func changeVal(iPtr *int) {
    *iPtr = 21
}

func main() {
  var i int = 11
  fmt.Println("value of i:",i)
  
  noChangeVal(i)
  fmt.Println("Unchanged value:",i)
  
  changeVal(&i)
  fmt.Println("Changed value:",i)
  
  noChangeVal(i)
  fmt.Println("Unchanged value:",i)
  
  iPtr := new(int)
  *iPtr = 50
  fmt.Println("value of iPtr:",*iPtr)
  
  changeVal(iPtr)
  fmt.Println("Changed value of iPtr:",*iPtr)
  
}