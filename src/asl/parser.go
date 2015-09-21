package asl

import (
    
)

const TAB = "    "

var tokens []Token
var tokenIndex int
var out string
var offset int

func Parse(token []Token) string {
    initParser(token)
    
    for tokenIndex < len(token) {
        parseBlock()
    }
    
    return out
}

// parser functions

func parseBlock() {
    if get().token == "var" {
        parseVar()
    } else if get().token == "if" {
        parseIf()
    } else {
        parseStatement()
    }
}

func parseVar() {
    expect("var")
    appendOut(get().token)
    next()
    
    if accept("=") {
        next()
        appendOut(" = "+get().token)
        next()
    }
    
    appendOut(";\n")
    expect(";")
}

func parseIf() {
    expect("if")
    appendOut("if (")
    parseCondition()
    appendOut(") then {\n")
    expect("{")
    parseBlock()
    expect("}")
    
    if accept("else") {
        next()
        expect("{")
        appendOut("} else {\n")
        parseBlock()
        expect("}")
    }
    
    appendOut("};")
}

func parseCondition() {
    for get().token != "{" {
        appendOut(get().token)
        next()
        
        if get().token != "{" {
            appendOut(" ")
        }
    }
}

func parseStatement() {
    
}

// helper functions

func initParser(token []Token) {
    if len(token) == 0 {
        panic("No tokens provided")
    }
    
    tokens = token
    tokenIndex = 0
    out = ""
    offset = 0
}

func accept(token string) bool {
    return tokenEqual(token, get())
}

func expect(token string) {
    if !tokenEqual(token, get()) {
        panic("Parse error, expected '"+token+"' but was '"+get().token+"'")
    }
    
    next()
}

func next() {
    tokenIndex++
}

func get() Token {
    if tokenIndex >= len(tokens) {
        panic("No more tokens")
    }
    
    return tokens[tokenIndex]
}

func tokenEqual(a string, b Token) bool {
    return a == b.token
}

func appendOut(str string) {
    out += str
}
