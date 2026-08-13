package main

import (
	"fmt"
	"time"
	)
	
func main() {
    messagesCh := make(chan string, 2)
    
    messagesCh <- "Buffered"
    messagesCh <- "Channel"
    
    go func(){
        messagesCh <- "Go"
    }()
    
    msg1 := <- messagesCh
    msg2 := <- messagesCh
    fmt.Println(msg1)
    fmt.Println(msg2)
    fmt.Println(<-messagesCh)
    time.Sleep(time.Second)
}
