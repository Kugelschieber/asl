package tokenizer

import (
	"strings"
)

type Token struct {
	Token        string
	Preprocessor bool
	Line         int
	Column       int
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
	"foreach",
	"func",
	"true",
	"false",
	"case",
	"default",
	"return",
	"try",
	"catch",
	"exitwith",
	"waituntil",
	"code"}

var whitespace = []byte{' ', '\n', '\t', '\r'}
var identifier = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
var preprocessor = byte('#')
var new_line = []byte{'\r', '\n'}

// Tokenizes the given byte array into syntax tokens,
// which can be parsed later.
func Tokenize(code []byte, doStripSlashes bool) []Token {
	if doStripSlashes {
		code = stripSlashes(code)
	}

	code = removeComments(code)
	tokens := make([]Token, 0)
	token, mask, isstring, line, column := "", false, false, 0, 0

	for i := 0; i < len(code); i++ {
		c := code[i]
		column++

		if byteArrayContains(new_line, c) {
			line++
			column = 0
		}

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
			// preprocessor, delimeter, keyword or variable/expression
			if c == preprocessor {
				tokens = append(tokens, preprocessorLine(code, &i, line, column))
				token = ""
			} else if byteArrayContains(delimiter, c) {
				if token != "" {
					tokens = append(tokens, Token{token, false, line, column})
				}

				tokens = append(tokens, Token{string(c), false, line, column})
				token = ""
			} else if stringArrayContains(strings.ToLower(token)) && !isIdentifierCharacter(c) {
				tokens = append(tokens, Token{token, false, line, column})
				token = ""
			} else if !byteArrayContains(whitespace, c) {
				token += string(c)
			}
		}

		mask = false
	}

	return tokens
}

// Removes slashes from input code.
// This is used for the "code" keyword for correct strings in resulting code.
func stripSlashes(code []byte) []byte {
	newcode := make([]byte, len(code))
	j, mask := 0, false

	for i := 0; i < len(code); i++ {
		c := code[i]

		if c == '\\' && !mask {
			mask = true
			continue
		}

		newcode[j] = code[i]
		mask = false
		j++
	}

	return newcode
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

// Reads preprocessor command until end of line
func preprocessorLine(code []byte, i *int, lineNr, column int) Token {
	c := byte('0')
	var line string

	for *i < len(code) {
		c = code[*i]

		if byteArrayContains(new_line, c) {
			break
		}

		line += string(c)
		(*i)++
	}

	// read all new line characters (\r and \n)
	c = code[*i]

	for byteArrayContains(new_line, c) {
		(*i)++
		c = code[*i]
	}

	(*i)-- // for will count up 1, so subtract it here

	return Token{line, true, lineNr, column}
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
func stringArrayContains(needle string) bool {
	for i := range keywords {
		if keywords[i] == needle {
			return true
		}
	}

	return false
}

// Checks if a character is allowed for identifiers.
func isIdentifierCharacter(c byte) bool {
	for i := range identifier {
		if identifier[i] == c {
			return true
		}
	}

	return false
}
