package asl

import (
    "strings"
    "fmt"
)

type Token struct{
    token string
}

var delimiter = []byte{
    '=',
    ';',
    '{',
    '}',
    '(',
    ')',
    '<',
    '>',
    '!',
    ',',
    ':',
    '&',
    '|',
    '$'}

var keywords = []string{
    "var",
    "if",
    "while",
    "switch",
    "for",
    "each",
    "func",
    "true",
    "false",
    "case",
    "default",
    "return",
    "sqfstart",
    "sqf"}

var whitespace = []byte{' ', '\n', '\t'}

func Tokenize(code []byte) []Token {
    code = removeComments(code)
    tokens := make([]Token, 0)
    token := ""
    
    fmt.Println(string(code))
    
    for i := range code {
        c := code[i]
        
        if byteArrayContains(delimiter, c) {
            if token != "" {
                tokens = append(tokens, Token{token})
            }
            
            tokens = append(tokens, Token{string(c)})
            token = ""
        } else if stringArrayContains(keywords, strings.ToLower(token)) {
            tokens = append(tokens, Token{token})
            token = ""
        } else if !byteArrayContains(whitespace, c) {
            token += string(c)
        }
    }
    
    return tokens
}

func removeComments(code []byte) []byte {
    newcode := make([]byte, len(code))
    j := 0
    
    for i := 0; i < len(code); i++ {
        c := code[i]
        
        if c == '/' && nextChar(code, i) == '/' {
            i = skipSingleLineComment(code, i+1)
            continue
        } else if c == '/' && nextChar(code, i) == '*' {
            i = skipMultiLineComment(code, i+1)
            continue
        }

        newcode[j] = c
        j++
    }
    
    return newcode[:j]
}

func nextChar(code []byte, i int) byte {
    i++
    
    if i < len(code) {
        return code[i]
    }
    
    return '0'
}

func skipSingleLineComment(code []byte, i int) int {
    for i < len(code) && code[i] != '\n' {
        i++
    }
    
    return i
}

func skipMultiLineComment(code []byte, i int) int {
    for i < len(code) && !(code[i] == '*' && nextChar(code, i) == '/') {
        i++
    }
    
    return i+1
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
