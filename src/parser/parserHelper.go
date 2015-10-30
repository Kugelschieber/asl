package parser

import (
    "tokenizer"
)

type Compiler struct {
    tokens []tokenizer.Token
    tokenIndex int
    out string
    offset int
    pretty bool
}

// Initilizes the parser.
func (c *Compiler) initParser(token []tokenizer.Token, prettyPrinting bool) bool {
	if len(token) == 0 {
		return false
	}

	c.tokens = token
	c.tokenIndex = 0
	c.out = ""
	c.offset = 0
	c.pretty = prettyPrinting
	
	return true
}

// Returns true, if current token matches expected one.
// Does not throw parse errors and checks if token is available.
func (c *Compiler) accept(token string) bool {
	return c.tokenIndex < len(c.tokens) && c.tokenEqual(token, c.get())
}

// Hard version of "accept".
// Throws if current token does not match expected one.
func (c *Compiler) expect(token string) {
	if !c.tokenEqual(token, c.get()) {
		panic("Parse error, expected '" + token + "' but was '" + c.get().Token + "'")
	}

	c.next()
}

// Returns true, if the next token matches expected one.
// Does not throw parse errors and checks if token is available.
func (c *Compiler) seek(token string) bool {
	if c.tokenIndex+1 >= len(c.tokens) {
		return false
	}

	return c.tokenEqual(token, c.tokens[c.tokenIndex+1])
}

// Increases token counter, so that the next token is compared.
func (c *Compiler) next() {
	c.tokenIndex++
}

// Returns current token or throws, if no more tokens are available.
func (c *Compiler) get() tokenizer.Token {
	if c.tokenIndex >= len(c.tokens) {
		panic("No more tokens")
	}

	return c.tokens[c.tokenIndex]
}

// Returns true if the end of input code was reached.
func (c *Compiler) end() bool {
	return c.tokenIndex == len(c.tokens)
}

// Checks if two strings match.
func (c *Compiler) tokenEqual(a string, b tokenizer.Token) bool {
	return a == b.Token
}

// Appends the output string to current SQF code output.
func (c *Compiler) appendOut(str string, newLine bool) {
	c.out += str

	if newLine && c.pretty {
		c.out += "\r\n"
	}
}
