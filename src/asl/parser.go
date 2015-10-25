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
	} else if accept("foreach") {
		parseForeach()
	} else if accept("func") {
		parseFunction()
	} else if accept("return") {
		parseReturn()
	} else if accept("case") || accept("default") {
		return
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

		if accept("[") {
			parseArray()
		} else {
			parseExpression(true)
		}
	}

	expect(";")
	appendOut(";", true)
}

func parseArray() {
	expect("[")
	appendOut("[", false)

	if !accept("]") {
		parseExpression(true)

		for accept(",") {
			next()
			appendOut(",", false)
			parseExpression(true)
		}
	}

	expect("]")
	appendOut("]", false)
}

func parseIf() {
	expect("if")
	appendOut("if (", false)
	parseExpression(true)
	appendOut(") then {", true)
	expect("{")
	parseBlock()
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
		next()
		appendOut("case ", false)
		parseExpression(true)
		expect(":")
		appendOut(":", true)

		if !accept("case") && !accept("}") && !accept("default") {
			appendOut("{", true)
			parseBlock()
			appendOut("};", true)
		}
	} else if accept("default") {
		next()
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
	expect("foreach")
	expr := parseExpression(false)
	expect("{")
	appendOut("{", true)
	parseBlock()
	expect("}")
	appendOut("} forEach ("+expr+");", true)
}

func parseFunction() {
	expect("func")
	appendOut(get().token+" = {", true)
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
		appendOut(name+" = _this select "+strconv.FormatInt(i, 10)+";", true)
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
	} else {
		parseFunctionCall(true, name)
		expect(";")
		appendOut(";", true)
	}

	if !end() {
		parseBlock()
	}
}

func parseAssignment() {
	expect("=")
	appendOut(" = ", false)
	parseExpression(true)
	expect(";")
	appendOut(";", true)
}

func parseFunctionCall(out bool, name string) string {
	output := ""

	expect("(")
	leftParams, leftParamCount := parseParameter(false)
	expect(")")

	if accept("(") {
		// buildin function
		next()
		rightParams, rightParamCount := parseParameter(false)
		expect(")")

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
		appendOut(output, false)
	}

	return output
}

func parseParameter(out bool) (string, int) {
	output := ""
	count := 0

	for !accept(")") {
		output += parseExpression(out)
		count++

		if !accept(")") {
			expect(",")
			output += ", "
		}
	}

	if out {
		appendOut(output, false)
	}

	return output, count
}

func parseExpression(out bool) string {
	output := parseArith()

	for accept("<") || accept(">") || accept("&") || accept("|") || accept("=") || accept("!") {
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
		} else if accept("=") {
			output += "="
			next()
		} else {
		    next()
		    expect("=")
		    output += "!="
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

	if seek("(") && !accept("!") && !accept("-") {
		name := get().token
		next()
		output = "(" + parseFunctionCall(false, name) + ")"
	} else if accept("!") || accept("-") {
		output = get().token
		next()

		if !accept("(") {
			output += get().token
			next()
		} else {
			output += parseTerm()
		}
	} else {
		output = get().token
		next()
	}

	return output
}

func parseTerm() string {
	if accept("(") {
		expect("(")
		output := "(" + parseExpression(false) + ")"
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
