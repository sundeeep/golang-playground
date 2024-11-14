package main

import (
	"fmt"
)

type BigStruct struct {
	buf [1 << 16]byte
}

var obj BigStruct = BigStruct{}

func funcPBV(obj BigStruct)  {}
func funcPBP(obj *BigStruct) {}

func main() {
	fmt.Println("Hello, World!")
	funcPBV(obj)
	funcPBV(obj)
}
