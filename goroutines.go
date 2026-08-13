package main

import (
	"fmt"
	"time"
)

func PrintName(name string) {
	for i := range 3 {
		fmt.Println(i+1, ":", name)
	}
}

func main() {
	PrintName("Ganesh")

	go PrintName("Jay")

	go func(name string) {
		fmt.Println(name)
	}("Anonymus")

	time.Sleep(time.Second)
	fmt.Println("Done...")
}

