package asl

import (
	"strconv"
)

const TAB = "    "

func Parse(token []Token) string {
	initParser(token)

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
	appendOut(get().token)
	next()

	if accept("=") {
		next()
		appendOut(" = ")
		parseExpression(true)
	}

	appendOut(";\n")
	expect(";")
}

func parseIf() {
	expect("if")
	appendOut("if (")
	parseExpression(true)
	appendOut(") then {\n")
	expect("{")
	parseBlock()
	expect("}")

	if accept("else") {
		next()
		expect("{")
		appendOut("} else {\n")
		parseBlock()
		expect("}")
	}

	appendOut("};\n")
}

func parseWhile() {
	expect("while")
	appendOut("while {")
	parseExpression(true)
	appendOut("} do {\n")
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};\n")
}

func parseSwitch() {
	expect("switch")
	appendOut("switch (")
	parseExpression(true)
	appendOut(") do {\n")
	expect("{")
	parseSwitchBlock()
	expect("}")
	appendOut("};\n")
}

func parseSwitchBlock() {
	if accept("}") {
		return
	}

	if accept("case") {
		expect("case")
		appendOut("case ")
		parseExpression(true)
		expect(":")
		appendOut(":\n")

		if !accept("case") && !accept("}") {
			appendOut("{\n")
			parseBlock()
			appendOut("};\n")
		}
	} else if accept("default") {
		expect("default")
		expect(":")
		appendOut("default:\n")

		if !accept("}") {
			appendOut("{\n")
			parseBlock()
			appendOut("};\n")
		}
	}

	parseSwitchBlock()
}

func parseFor() {
	expect("for")
	appendOut("for [{")

	// var in first assignment is optional
	if accept("var") {
		next()
	}

	parseExpression(true)
	expect(";")
	appendOut("}, {")
	parseExpression(true)
	expect(";")
	appendOut("}, {")
	parseExpression(true)
	expect(";")
	appendOut("}] do {\n")
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};\n")
}

func parseForeach() {
	expect("each")
	expr := parseExpression(false)
	expect("{")
	appendOut("{\n")
	parseBlock()
	expect("}")
	appendOut("} forEach (" + expr + ");\n")
}

func parseFunction() {
	expect("func")
	appendOut(get().token + " = {\n")
	next()
	expect("(")
	parseFunctionParameter()
	expect(")")
	expect("{")
	parseBlock()
	expect("}")
	appendOut("};\n")
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
		appendOut(name + " = _this select " + strconv.FormatInt(i, 10) + ";\n")
		i++

		if !accept(")") {
			expect(",")
		}
	}
}

func parseReturn() {
	expect("return")
	appendOut("return ")
	parseExpression(true)
	expect(";")
	appendOut(";\n")
}

func parseSqf() {
	expect("sqf")
	expect(":")

	for !accept("sqf") {
		appendOut(get().token)
		next()
	}

	appendOut("\n")
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
		appendOut(name)
		parseAssignment()
	} else if name == "$" {
		name = get().token
		next()
		parseBuildinFunctionCall(name)
	} else {
		parseFunctionCall()
		appendOut(name + ";\n")
	}

	if !end() {
		parseBlock()
	}
}

func parseAssignment() {
	expect("=")
	appendOut(" = " + get().token)
	next()
	expect(";")
	appendOut(";\n")
}

func parseFunctionCall() {
	expect("(")
	appendOut("[")
	parseParameter()
	expect(")")
	expect(";")
	appendOut("] call ")
}

func parseBuildinFunctionCall(name string) {
	expect("(")
	appendOut("[")
	parseParameter()
	expect(")")
	appendOut("] ")
	expect("(")
	appendOut(name + " [")
	parseParameter()
	expect(")")
	expect(";")
	appendOut("];\n")
}

func parseParameter() {
	for !accept(")") {
		parseExpression(true)

		if !accept(")") {
			expect(",")
			appendOut(", ")
		}
	}
}

func parseExpression(out bool) string {
	openingBrackets := 0
	output := ""

	for !accept(",") && !accept(":") && !accept(";") && !accept("{") && !accept("}") && (openingBrackets != 0 || !accept(")")) {
		current := get().token

		if out {
			appendOut(current)
		} else {
			output += current
		}

		if accept("(") {
			openingBrackets++
		} else if accept(")") {
			openingBrackets--
		}

		next()
	}

	return output
}
