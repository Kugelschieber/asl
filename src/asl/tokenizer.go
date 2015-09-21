package asl

import (
    "fmt"
)

type Token struct{
    token string
}

var delimiter = []byte{'=',
    ';',
    '{',
    '}',
    '<',
    '>',
    '!'}

var keywords = []string{"var",
    "if"}

var whitespace = []byte{' ', '\n', '\t'}

func Tokenize(code []byte) []Token {
    tokens := make([]Token, 0)
    token := ""
    
    for i := range code {
        c := code[i]
        
        if byteArrayContains(delimiter, c) {
            if token != "" {
                tokens = append(tokens, Token{token})
            }
            
            tokens = append(tokens, Token{string(c)})
            token = ""
        } else if stringArrayContains(keywords, token) {
            tokens = append(tokens, Token{token})
            token = ""
        } else if !byteArrayContains(whitespace, c) {
            token += string(c)
        }
    }
    
    // TEST
    fmt.Println("Tokens:")
    for i := range tokens {
        fmt.Println(tokens[i].token)
    }
    fmt.Println("---")
    
    return tokens
}

func byteArrayContains(haystack []byte, needle byte) bool {
    for i := range haystack {
        if haystack[i] == needle {
            return true;
        }
    }
    
    return false
}

func stringArrayContains(haystack []string, needle string) bool {
    for i := range haystack {
        if haystack[i] == needle {
            return true;
        }
    }
    
    return false
}
