package main

import (
    "io/ioutil"
    "asl"
)

func main(){
    // read test file
    code, _ := ioutil.ReadFile("in/simple.asl")
    asl.Tokenize(code)
}
