package parser

import (
	"tokenizer"
)

// Parses tokens, validates code to a specific degree
// and writes SQF code into desired location.
func (c *Compiler) Parse(token []tokenizer.Token, prettyPrinting bool) string {
	if !c.initParser(token, prettyPrinting) {
	    return ""
	}

	for c.tokenIndex < len(token) {
		c.parseBlock()
	}

	return c.out
}

func (c *Compiler) parseBlock() {
	if c.accept("var") {
		c.parseVar()
	} else if c.accept("if") {
		c.parseIf()
	} else if c.accept("while") {
		c.parseWhile()
	} else if c.accept("switch") {
		c.parseSwitch()
	} else if c.accept("for") {
		c.parseFor()
	} else if c.accept("foreach") {
		c.parseForeach()
	} else if c.accept("func") {
		c.parseFunction()
	} else if c.accept("return") {
		c.parseReturn()
	} else if c.accept("try") {
		c.parseTryCatch()
	} else if c.accept("exitwith") {
	    c.parseExitWith()
	} else if c.accept("waituntil") {
	    c.parseWaitUntil()
	} else if c.accept("case") || c.accept("default") {
		return
	} else {
		c.parseStatement()
	}

	if !c.end() && !c.accept("}") {
		c.parseBlock()
	}
}

func (c *Compiler) parseVar() {
	c.expect("var")
	c.appendOut(c.get().Token, false)
	c.next()

	if c.accept("=") {
		c.next()
		c.appendOut(" = ", false)

		if c.accept("[") {
			c.parseArray()
		} else {
			c.parseExpression(true)
		}
	}

	c.expect(";")
	c.appendOut(";", true)
}

func (c *Compiler) parseArray() {
	c.expect("[")
	c.appendOut("[", false)

	if !c.accept("]") {
		c.parseExpression(true)

		for c.accept(",") {
			c.next()
			c.appendOut(",", false)
			c.parseExpression(true)
		}
	}

	c.expect("]")
	c.appendOut("]", false)
}

func (c *Compiler) parseIf() {
	c.expect("if")
	c.appendOut("if (", false)
	c.parseExpression(true)
	c.appendOut(") then {", true)
	c.expect("{")
	c.parseBlock()
	c.expect("}")

	if c.accept("else") {
		c.next()
		c.expect("{")
		c.appendOut("} else {", true)
		c.parseBlock()
		c.expect("}")
	}

	c.appendOut("};", true)
}

func (c *Compiler) parseWhile() {
	c.expect("while")
	c.appendOut("while {", false)
	c.parseExpression(true)
	c.appendOut("} do {", true)
	c.expect("{")
	c.parseBlock()
	c.expect("}")
	c.appendOut("};", false)
}

func (c *Compiler) parseSwitch() {
	c.expect("switch")
	c.appendOut("switch (", false)
	c.parseExpression(true)
	c.appendOut(") do {", true)
	c.expect("{")
	c.parseSwitchBlock()
	c.expect("}")
	c.appendOut("};", true)
}

func (c *Compiler) parseSwitchBlock() {
	if c.accept("}") {
		return
	}

	if c.accept("case") {
		c.next()
		c.appendOut("case ", false)
		c.parseExpression(true)
		c.expect(":")
		c.appendOut(":", true)

		if !c.accept("case") && !c.accept("}") && !c.accept("default") {
			c.appendOut("{", true)
			c.parseBlock()
			c.appendOut("};", true)
		}
	} else if c.accept("default") {
		c.next()
		c.expect(":")
		c.appendOut("default:", true)

		if !c.accept("}") {
			c.appendOut("{", true)
			c.parseBlock()
			c.appendOut("};", true)
		}
	}

	c.parseSwitchBlock()
}

func (c *Compiler) parseFor() {
	c.expect("for")
	c.appendOut("for [{", false)

	// var in first assignment is optional
	if c.accept("var") {
		c.next()
	}

	c.parseExpression(true)
	c.expect(";")
	c.appendOut("}, {", false)
	c.parseExpression(true)
	c.expect(";")
	c.appendOut("}, {", false)
	c.parseExpression(true)
	c.appendOut("}] do {", true)
	c.expect("{")
	c.parseBlock()
	c.expect("}")
	c.appendOut("};", true)
}

func (c *Compiler) parseForeach() {
	c.expect("foreach")
	expr := c.parseExpression(false)
	c.expect("{")
	c.appendOut("{", true)
	c.parseBlock()
	c.expect("}")
	c.appendOut("} forEach ("+expr+");", true)
}

func (c *Compiler) parseFunction() {
	c.expect("func")
	c.appendOut(c.get().Token+" = {", true)
	c.next()
	c.expect("(")
	c.parseFunctionParameter()
	c.expect(")")
	c.expect("{")
	c.parseBlock()
	c.expect("}")
	c.appendOut("};", true)
}

func (c *Compiler) parseFunctionParameter() {
	// empty parameter list
	if c.accept("{") {
		return
	}
	
	c.appendOut("params [", false)

	for !c.accept(")") {
		name := c.get().Token
		c.next()
		
		if c.accept("=") {
		    c.next()
		    value := c.get().Token
		    c.next()
		    c.appendOut("[\""+name+"\","+value+"]", false)
		} else {
		    c.appendOut("\""+name+"\"", false)
		}

		if !c.accept(")") {
			c.expect(",")
			c.appendOut(",", false)
		}
	}
	
	c.appendOut("];", true)
}

func (c *Compiler) parseReturn() {
	c.expect("return")
	c.appendOut("return ", false)
	c.parseExpression(true)
	c.expect(";")
	c.appendOut(";", true)
}

func (c *Compiler) parseTryCatch() {
	c.expect("try")
	c.expect("{")
	c.appendOut("try {", true)
	c.parseBlock()
	c.expect("}")
	c.expect("catch")
	c.expect("{")
	c.appendOut("} catch {", true)
	c.parseBlock()
	c.expect("}")
	c.appendOut("};", true)
}

func (c *Compiler) parseExitWith() {
    c.expect("exitwith")
    c.expect("{")
    c.appendOut("if (true) exitWith {", true)
    c.parseBlock()
    c.expect("}")
    c.appendOut("};", true)
}

func (c *Compiler) parseWaitUntil() {
    c.expect("waituntil")
    c.expect("(")
    c.appendOut("waitUntil {", false)
    c.parseExpression(true)
    
    if c.accept(";") {
        c.next()
        c.appendOut(";", false)
        c.parseExpression(true)
    }
    
    c.expect(")")
    c.expect(";")
    c.appendOut("};", true)
}

func (c *Compiler) parseInlineCode() string {
    c.expect("code")
    c.expect("(")
    
    code := c.get().Token
    c.next()
    output := "{}"
    
    if len(code) > 2 {
        compiler := Compiler{}
        output = "{"+compiler.Parse(tokenizer.Tokenize([]byte(code[1:len(code)-1])), false)+"}"
    }
    
    c.expect(")")
    
    return output
}

// Everything that does not start with a keyword.
func (c *Compiler) parseStatement() {
	// empty block
	if c.accept("}") || c.accept("case") || c.accept("default") {
		return
	}

	// variable or function name
	name := c.get().Token
	c.next()

	if c.accept("=") {
		c.appendOut(name, false)
		c.parseAssignment()
	} else {
		c.parseFunctionCall(true, name)
		c.expect(";")
		c.appendOut(";", true)
	}

	if !c.end() {
		c.parseBlock()
	}
}

func (c *Compiler) parseAssignment() {
	c.expect("=")
	c.appendOut(" = ", false)
	c.parseExpression(true)
	c.expect(";")
	c.appendOut(";", true)
}

func (c *Compiler) parseFunctionCall(out bool, name string) string {
	output := ""

	c.expect("(")
	leftParams, leftParamCount := c.parseParameter(false)
	c.expect(")")

	if c.accept("(") {
		// buildin function
		c.next()
		rightParams, rightParamCount := c.parseParameter(false)
		c.expect(")")

		if leftParamCount > 1 {
			leftParams = "[" + leftParams + "]"
		}

		if rightParamCount > 1 {
			rightParams = "[" + rightParams + "]"
		}

		if leftParamCount > 0 {
			output = leftParams + " " + name + " " + rightParams
		} else {
			output = name + " " + rightParams
		}
	} else {
		output = "[" + leftParams + "] call " + name
	}

	if out {
		c.appendOut(output, false)
	}

	return output
}

func (c *Compiler) parseParameter(out bool) (string, int) {
	output := ""
	count := 0

	for !c.accept(")") {
		output += c.parseExpression(out)
		count++

		if !c.accept(")") {
			c.expect(",")
			output += ", "
		}
	}

	if out {
		c.appendOut(output, false)
	}

	return output, count
}

func (c *Compiler) parseExpression(out bool) string {
	output := c.parseArith()

	for c.accept("<") || c.accept(">") || c.accept("&") || c.accept("|") || c.accept("=") || c.accept("!") {
		if c.accept("<") {
			output += "<"
			c.next()
		} else if c.accept(">") {
			output += ">"
			c.next()
		} else if c.accept("&") {
			c.next()
			c.expect("&")
			output += "&&"
		} else if c.accept("|") {
			c.next()
			c.expect("|")
			output += "||"
		} else if c.accept("=") {
			output += "="
			c.next()
		} else {
			c.next()
			c.expect("=")
			output += "!="
		}

		if c.accept("=") {
			output += "="
			c.next()
		}

		output += c.parseExpression(false)
	}

	if out {
		c.appendOut(output, false)
	}

	return output
}

func (c *Compiler) parseIdentifier() string {
	output := ""

	if c.accept("code") {
	    output += c.parseInlineCode()
	} else if c.seek("(") && !c.accept("!") && !c.accept("-") {
		name := c.get().Token
		c.next()
		output = "(" + c.parseFunctionCall(false, name) + ")"
	} else if c.seek("[") {
	    output += "("+c.get().Token
	    c.next()
	    c.expect("[")
	    output += " select ("+c.parseExpression(false)+"))"
	    c.expect("]")
	} else if c.accept("!") || c.accept("-") {
		output = c.get().Token
		c.next()
		output += c.parseTerm()
	} else {
		output = c.get().Token
		c.next()
	}

	return output
}

func (c *Compiler) parseTerm() string {
	if c.accept("(") {
		c.expect("(")
		output := "(" + c.parseExpression(false) + ")"
		c.expect(")")

		return output
	}

	return c.parseIdentifier()
}

func (c *Compiler) parseFactor() string {
	output := c.parseTerm()

	for c.accept("*") || c.accept("/") { // TODO: modulo?
		if c.accept("*") {
			output += "*"
		} else {
			output += "/"
		}

		c.next()
		output += c.parseExpression(false)
	}

	return output
}

func (c *Compiler) parseArith() string {
	output := c.parseFactor()

	for c.accept("+") || c.accept("-") {
		if c.accept("+") {
			output += "+"
		} else {
			output += "-"
		}

		c.next()
		output += c.parseExpression(false)
	}

	return output
}
