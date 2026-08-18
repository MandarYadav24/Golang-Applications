package main
import (
  "fmt"
)

func main() {
  queues := make(chan string, 2)
  queues <- "one"
  queues <- "two"

  close(queues)

  for ele := range queues{
    fmt.Println(ele)
  }

}