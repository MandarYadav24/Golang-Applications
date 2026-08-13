package main

import (
	"fmt"
	"time"
	)
	
func Worker(done chan bool) {
    fmt.Println("Worker start...")
    time.Sleep(time.Second)
    fmt.Println("Work done.")
    
    done <-true
}
	
func main() {
    done := make(chan bool, 1)
    
    go Worker(done)
    
    <-done
}
