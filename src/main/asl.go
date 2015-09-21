package main

import (
	"asl"
	"fmt"
	"io/ioutil"
)

func main() {
	// read test file
	code, _ := ioutil.ReadFile("in/simple.asl")
	token := asl.Tokenize(code)
	out := asl.Parse(token)

	fmt.Println(out)
}
