package main

import (
	"asl"
	"fmt"
	"io/ioutil"
)

func main() {
	// read test file
	code, _ := ioutil.ReadFile("in/test.asl")
	token := asl.Tokenize(code)
	out := asl.Parse(token)

	fmt.Println("OUTPUT:\n"+out) // TODO: remove
}
