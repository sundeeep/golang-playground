## Go Concurrency Paterns.

Write performant and maintainable code in go using concurrency patterns.

### Go Concurrency Primitives
- Go Routines
- Channels
- Select

#### Go Routines
- Example of using go routines concept, practically.
```go
package main

import (
	"fmt"
)

func someFunc(num int) {
	fmt.Println(num)
}

// main function is the entry point to the go process.
func main() {
	go someFunc(1) // Go routine
	go someFunc(1) // Go routine
	go someFunc(1) // Go routine
	go someFunc(1) // Go routine
	fmt.Println("hi")
}
```