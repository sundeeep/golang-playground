package main

import (
	"fmt"
	"time"
)

func someFunc(num int) {
	fmt.Println(num)
}

// main function is the entry point to the go process.
func main() {

	// Below three functions will get invoked concurrently in no specific order.
	go someFunc(1) // Go routine
	go someFunc(2) // Go routine
	go someFunc(3) // Go routine

	time.Sleep(time.Second * 2)
	fmt.Println("hi")
}
