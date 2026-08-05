package main
import (
    "fmt"
    "errors"
    )

type argError struct {
    arg int
    message string
}

func (e *argError)Error() string{
    return fmt.Sprintf("%d - %s",e.arg, e.message)
}

func checkError(arg int) (int, error) {
    if arg == 101 {
        return -1, &argError{arg, "Invalid ID"}
    } 
    return arg, nil
}

func main() {
  e := argError{11,"Invalid Id"}
  fmt.Println(e.Error())
  
  _, err := checkError(101)
  if ae, ok := errors.AsType[*argError](err); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.message)
    } else {
        fmt.Println("err doesn't match argError")
    }
}