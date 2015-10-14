package asl

import (
	"strconv"
)

const TAB = "    "

// Parses tokens, validates code to a specific degree
// and writes SQF code into desired location.
func Parse(token []Token, prettyPrinting bool) string {
	initParser(token, prettyPrinting)

	for tokenIndex < len(token) {
		parseBlock()
	}

	return out
}

func parseBlock() {
	if accept("var") {
		parseVar()
	} else if accept("if") {
		parseIf()
	} else if accept("while") {
		parseWhile()
	} else if accept("switch") {
		parseSwitch()
	} else if accept("for") {
		parseFor()
	} else if accept("each") {
		parseForeach()
	} else if accept("func") {
		parseFunction()
	} else if accept("return") {
		parseReturn()
	} else if accept("sqf") {
		parseSqf()
	} else {
		parseStatement()
	}

	if !end() && !accept("}") {
		parseBlock()
	}
}

func parseVar() {
	expect("var")
	appendOut(get().token, false)
	next()

	if accept("=") {
		next()
		appendOut(" = ", false)
		parseExpression(true)
	}

    expect(";")
	appendOut(";", true)
}

func parseIf() {
	expect("if")
	appendOut("if (", false)
	parseExpression(true)
	appendOut(") then {", true)
	expect("{")
	parseExpression(true)
	expect("}")

	if accept("else") {
		next()
		expect("{")
		appendOut("} else {", true)
		parseBlock()
		expect("}")
	}

	appendOut("};", true)
}

func parseWhile() {
	expect("while")
	appendOut("while {", false)
	parseExpression(true)
	appendOut("} do {", true)
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};", false)
}

func parseSwitch() {
	expect("switch")
	appendOut("switch (", false)
	parseExpression(true)
	appendOut(") do {", true)
	expect("{")
	parseSwitchBlock()
	expect("}")
	appendOut("};", true)
}

func parseSwitchBlock() {
	if accept("}") {
		return
	}

	if accept("case") {
		expect("case")
		appendOut("case ", false)
		parseExpression(true)
		expect(":")
		appendOut(":", true)

		if !accept("case") && !accept("}") {
			appendOut("{", true)
			parseBlock()
			appendOut("};", true)
		}
	} else if accept("default") {
		expect("default")
		expect(":")
		appendOut("default:", true)

		if !accept("}") {
			appendOut("{", true)
			parseBlock()
			appendOut("};", true)
		}
	}

	parseSwitchBlock()
}

func parseFor() {
	expect("for")
	appendOut("for [{", false)

	// var in first assignment is optional
	if accept("var") {
		next()
	}

	parseExpression(true)
	expect(";")
	appendOut("}, {", false)
	parseExpression(true)
	expect(";")
	appendOut("}, {", false)
	parseExpression(true)
	appendOut("}] do {", true)
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};", true)
}

func parseForeach() {
	expect("each")
	expr := parseExpression(false)
	expect("{")
	appendOut("{", true)
	parseBlock()
	expect("}")
	appendOut("} forEach (" + expr + ");", true)
}

func parseFunction() {
	expect("func")
	appendOut(get().token + " = {", true)
	next()
	expect("(")
	parseFunctionParameter()
	expect(")")
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};", true)
}

func parseFunctionParameter() {
	// empty parameter list
	if accept("{") {
		return
	}

	i := int64(0)

	for !accept(")") {
		name := get().token
		next()
		appendOut(name + " = _this select " + strconv.FormatInt(i, 10) + ";", true)
		i++

		if !accept(")") {
			expect(",")
		}
	}
}

func parseReturn() {
	expect("return")
	appendOut("return ", false)
	parseExpression(true)
	expect(";")
	appendOut(";", true)
}

func parseSqf() {
	expect("sqf")
	expect(":")

	for !accept("sqf") {
		appendOut(get().token, false)
		next()
	}

	appendOut("", true)
	expect("sqf")
}

// Everything that does not start with a keyword.
func parseStatement() {
	// empty block
	if accept("}") || accept("case") || accept("default") {
		return
	}

	// variable or function name
	name := get().token
	next()

	if accept("=") {
		appendOut(name, false)
		parseAssignment()
	} else if name == "$" {
		name = get().token
		next()
		parseBuildinFunctionCall(name)
	} else {
		parseFunctionCall(true)
		appendOut(name + ";", true)
	}

	if !end() {
		parseBlock()
	}
}

func parseAssignment() {
	expect("=")
	appendOut(" = " + get().token, false)
	next()
	expect(";")
	appendOut(";", true)
}

func parseFunctionCall(out bool) string {
    output := "["
    
	expect("(")
	//output += parseParameter()
	expect(")")
	//expect(";")
	output += "] call "
	
	if out {
	    appendOut(output, true)
	}
	
	return output
}

func parseBuildinFunctionCall(name string) {
    // FIXME: does not work for all kind of commands
	expect("(")
	appendOut("[", false)
	parseParameter()
	expect(")")
	appendOut("] ", false)
	expect("(")
	appendOut(name + " [", false)
	parseParameter()
	expect(")")
	expect(";")
	appendOut("];", true)
}

func parseParameter() {
	for !accept(")") {
		parseExpression(true)

		if !accept(")") {
			expect(",")
			appendOut(", ", false)
		}
	}
}

func parseExpression(out bool) string {
	output := parseArith()
	
	for accept("<") || accept(">") || accept("&") || accept("|") || accept("=") {
	    if accept("<") {
	        output += "<"
	        next()
	    } else if accept(">") {
	        output += ">"
	        next()
	    } else if accept("&") {
	        next()
	        expect("&")
	        output += "&&"
	    } else if accept("|") {
	        next()
	        expect("|")
	        output += "||"
	    } else {
	        output += "="
	        next()
	    }
	    
	    if accept("=") {
	        output += "="
	        next()
	    }
	    
	    output += parseExpression(false)
	}
	
	if out {
	    appendOut(output, false)
	}
	
	return output
}

func parseIdentifier() string {
    output := ""
    
    if seek("(") {
        name := get().token
        next()
        output = "("+parseFunctionCall(false)+name+")"
    } else {
        output = get().token
        next()
    }
    
    return output
}

func parseTerm() string {
    if accept("(") {
        expect("(")
        output := "("+parseExpression(false)+")"
        expect(")")
        
        return output
    }
    
    return parseIdentifier()
}

func parseFactor() string {
    output := parseTerm()
    
    for accept("*") || accept("/") { // TODO: modulo?
        if accept("*") {
            output += "*"
        } else {
            output += "/"
        }
        
        next()
        output += parseExpression(false)
    }
    
    return output
}

func parseArith() string {
    output := parseFactor()
	
    for accept("+") || accept("-") {
        if accept("+") {
            output += "+"
        } else {
            output += "-"
        }
        
        next()
        output += parseExpression(false)
    }
	
	return output
}
