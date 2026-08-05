package main
import (
    "fmt"
    "errors"
    )

func Flight(fltNo int) (int, error){
    if fltNo == 100 {
        return -1, errors.New("Flight no 100 is not available")
    }
    return fltNo, nil
}

var ErrNoFlt = errors.New("No flight available")
var ErrNoSeat = errors.New("No seats available")

func BookFlight(fltNo int) error {
    if fltNo == 100 {
        return ErrNoFlt
    } else if fltNo == 101 {
        return fmt.Errorf("Flight available but %w",ErrNoSeat)
    }
    return nil
}

func main() {
  for _, v := range []int{10,20,100,50,40}{
      if r, e := Flight(v); e != nil {
          fmt.Println("Flight failed:",e)
      } else {
          fmt.Println("Flight worked:",r)
      }
  }
  
  var flightNos = []int {102,100,101,200,300}
  
  for _, v := range flightNos {
      if err := BookFlight(v); err != nil {
          if errors.Is(err, ErrNoFlt) {
              fmt.Printf("We should book another flight:%v\n",v)
          } else if errors.Is(err, ErrNoSeat) {
              fmt.Printf("All seats booked:%v\n",v)
          } else {
              fmt.Printf("unknown error: %s\n",err)
          }
          continue
      }
      fmt.Println("Flight is booked successfully..")
  }
}