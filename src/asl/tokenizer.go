package asl

import (
	"strings"
)

type Token struct {
	token string
}

var delimiter = []byte{
	'=',
	';',
	'{',
	'}',
	'(',
	')',
	'[',
	']',
	'<',
	'>',
	'!',
	',',
	':',
	'&',
	'|',
	'+',
	'-',
	'*',
	'/'} // TODO: modulo?

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
	"return"}

var whitespace = []byte{' ', '\n', '\t'}

// Tokenizes the given byte array into syntax tokens,
// which can be parsed later.
func Tokenize(code []byte) []Token {
	code = removeComments(code)
	tokens := make([]Token, 0)
	token, mask, isstring := "", false, false

	for i := range code {
		c := code[i]
		
		// string masks (backslash)
		if c == '\\' && !mask {
		    token += "\\"
		    mask = true
		    continue
		}
		
		// string
		if c == '"' && !mask {
		    token += "\""
		    isstring = !isstring
		    continue
		}
		
		if isstring {
		    token += string(c)
		} else {
            // delimeter, keyword or variable/expression
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
		
		mask = false
	}

	return tokens
}

// Removes all comments from input byte array.
// Comments are single line comments, starting with // (two slashes),
// multi line comments with /* ... */ (slash star, star slash).
func removeComments(code []byte) []byte {
	newcode := make([]byte, len(code))
	j, mask, isstring := 0, false, false

	for i := 0; i < len(code); i++ {
		c := code[i]
		
		// do not remove comments from strings
		if c == '\\' && !mask {
		    mask = true
		}
		
		if c == '"' && !mask {
		    isstring = !isstring
		}

        // single/multi line comment
        if !isstring {
    		if c == '/' && nextChar(code, i) == '/' {
    			i = skipSingleLineComment(code, i+1)
    			continue
    		} else if c == '/' && nextChar(code, i) == '*' {
    			i = skipMultiLineComment(code, i+1)
    			continue
    		}
        }

		newcode[j] = c
		j++
		mask = false
	}

	return newcode[:j]
}

// Returns the next character in code starting at i.
// If no character is left, '0' will be returned.
func nextChar(code []byte, i int) byte {
	i++

	if i < len(code) {
		return code[i]
	}

	return '0'
}

// Used to skip a line if a single line comment was found.
func skipSingleLineComment(code []byte, i int) int {
	for i < len(code) && code[i] != '\n' {
		i++
	}

	return i
}

// Used to skip a block of characters if a multi line comment was found
func skipMultiLineComment(code []byte, i int) int {
	for i < len(code) && !(code[i] == '*' && nextChar(code, i) == '/') {
		i++
	}

	return i + 1
}

// Checks if a byte array (string) contains a delimeter.
func byteArrayContains(haystack []byte, needle byte) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}

// Checks if a byte array (string) contains a string delimeter.
func stringArrayContains(haystack []string, needle string) bool {
	for i := range haystack {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}
