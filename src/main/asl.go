package main

import (
	"asl"
	"fmt"
	"io/ioutil"
)

func usage() {
    fmt.Println("Usage: asl [-v|-pretty]")
}

func main() {
	// read test file
	code, _ := ioutil.ReadFile("in/simple.asl")
	token := asl.Tokenize(code)
	out := asl.Parse(token, true)

	fmt.Println("OUTPUT:\n"+out) // TODO: remove
}
