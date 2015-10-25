package asl

var tokens []Token
var tokenIndex int
var out string
var offset int
var pretty bool

// Initilizes the parser.
func initParser(token []Token, prettyPrinting bool) bool {
	if len(token) == 0 {
		return false
	}

	tokens = token
	tokenIndex = 0
	out = ""
	offset = 0
	pretty = prettyPrinting
	
	return true
}

// Returns true, if current token matches expected one.
// Does not throw parse errors and checks if token is available.
func accept(token string) bool {
	return tokenIndex < len(tokens) && tokenEqual(token, get())
}

// Hard version of "accept".
// Throws if current token does not match expected one.
func expect(token string) {
	if !tokenEqual(token, get()) {
		panic("Parse error, expected '" + token + "' but was '" + get().token + "'")
	}

	next()
}

// Returns true, if the next token matches expected one.
// Does not throw parse errors and checks if token is available.
func seek(token string) bool {
	if tokenIndex+1 >= len(tokens) {
		return false
	}

	return tokenEqual(token, tokens[tokenIndex+1])
}

// Increases token counter, so that the next token is compared.
func next() {
	tokenIndex++
}

// Returns current token or throws, if no more tokens are available.
func get() Token {
	if tokenIndex >= len(tokens) {
		panic("No more tokens")
	}

	return tokens[tokenIndex]
}

// Returns true if the end of input code was reached.
func end() bool {
	return tokenIndex == len(tokens)
}

// Checks if two strings match.
func tokenEqual(a string, b Token) bool {
	return a == b.token
}

// Appends the output string to current SQF code output.
func appendOut(str string, newLine bool) {
	out += str

	if newLine && pretty {
		out += "\n"
	}
}
