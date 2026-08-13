package main

import (
	"fmt"
	)
	
func main() {
    messagesCh := make(chan string)
    
    go func() {
      messagesCh <- "Ganesh"  
    }()
    
    msg := <- messagesCh
    fmt.Println(msg)
    
}
