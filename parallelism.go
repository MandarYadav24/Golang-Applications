package main

import (
  "fmt"
  "runtime"
  "sync"
)

func calculate(name string, wg *sync.WaitGroup) {
  defer wg.Done()
  sum := 0
  for i := 0; i <= 1000; i++{
    sum += i
  }
  fmt.Println(name,"completed:",sum)
}

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())

  var wg sync.WaitGroup

  wg.Add(2)
  go calculate("Task 1",&wg)
  go calculate("Task 2",&wg)

  wg.Wait()

  fmt.Println("Both task completed...")
}

