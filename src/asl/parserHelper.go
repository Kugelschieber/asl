package asl

var tokens []Token
var tokenIndex int
var out string
var offset int

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
    return tokenIndex < len(tokens) && tokenEqual(token, get())
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

func end() bool {
    return tokenIndex == len(tokens)
}

func tokenEqual(a string, b Token) bool {
    return a == b.token
}

func appendOut(str string) {
    out += str
}
