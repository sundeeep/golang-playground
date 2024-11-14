package main

type BigStruct struct {
	buf [1 << 16]byte
}

var obj BigStruct = BigStruct{}

func funcPBP(obj *BigStruct) {}

func funcPBV(obj BigStruct) {}

func main() {
	funcPBP(&obj)
	funcPBV(obj)
}
