package main

import (
    "io/ioutil"
    "asl"
    "fmt"
)

func main(){
    // read test file
    code, _ := ioutil.ReadFile("in/statements.asl")
    token := asl.Tokenize(code)
    out := asl.Parse(token)
    
    fmt.Println(out)
}
