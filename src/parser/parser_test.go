package parser_test

import (
	"io/ioutil"
	"parser"
	"testing"
	"tokenizer"
	"types"
)

const (
	types_file = "../../test/types"
)

func TestParserDeclaration(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_var.asl")
	want := "x = 1;\r\narray = [1,2,3];\r\n"

	equal(t, got, want)
}

func TestParserAssignment(t *testing.T) {
	got := getCompiled(t, "../../test/parser_assignment.asl")
	want := "x = 1;\r\n"

	equal(t, got, want)
}

func TestParserIf(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_if.asl")
	want := "if (a<b) then {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserWhile(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_while.asl")
	want := "while {true} do {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserFor(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_for.asl")
	want := "for [{i=0}, {i<100}, {i=i+1}] do {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserForeach(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_foreach.asl")
	want := "{\r\nunit = _x;\r\n} forEach (allUnits);\r\n"

	equal(t, got, want)
}

func TestParserSwitch(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_switch.asl")
	want := "switch (x) do {\r\ncase 1:\r\n{\r\nx = 1;\r\n};\r\ncase 2:\r\n{\r\nx = 2;\r\n};\r\ndefault:\r\n{\r\nx = 3;\r\n};\r\n};\r\n"

	equal(t, got, want)
}

func TestParserFunction(t *testing.T) {
	got := getCompiled(t, "../../test/tokenizer_func.asl")
	want := "TestFunction = {\r\nparams [\"param0\",\"param1\"];\r\nreturn true;\r\n};\r\n"

	equal(t, got, want)
}

func TestParserAssignResult(t *testing.T) {
	got := getCompiled(t, "../../test/parser_assign_result.asl")
	want := "x = ([1, 2, 3] call foo);\r\ny = ([1, 2, 3] call bar);\r\n"

	equal(t, got, want)
}

func TestParserExpression(t *testing.T) {
	got := getCompiled(t, "../../test/parser_expression.asl")
	want := "x = -(1+(2+3))/(6*(someVariable+99-100))-(20)+!anotherVariable+([] call foo);\r\n"

	equal(t, got, want)
}

func TestParserExpression2(t *testing.T) {
	got := getCompiled(t, "../../test/parser_expression2.asl")
	want := "x = true||(3>=4&&5<8);\r\n"

	equal(t, got, want)
}

func TestParserFunctionCall(t *testing.T) {
	got := getCompiled(t, "../../test/parser_func_call.asl")
	want := "myFunc = {\r\nparams [\"a\",\"b\"];\r\nreturn a>b;\r\n};\r\n[1+3/4, 2-(66*22)/3-((123))] call myFunc;\r\n"

	equal(t, got, want)
}

func TestParserNullBuildinFunctionCall(t *testing.T) {
	types.LoadTypes(types_file)

	got := getCompiled(t, "../../test/parser_null_buildin_func.asl")
	want := "_volume = (radioVolume);\r\n"

	equal(t, got, want)
}

func TestParserUnaryBuildinFunctionCall(t *testing.T) {
	types.LoadTypes(types_file)

	got := getCompiled(t, "../../test/parser_unary_buildin_func.asl")
	want := "_isReady = (unitReady soldier);\r\n"

	equal(t, got, want)
}

func TestParserBinaryBuildinFunctionCall(t *testing.T) {
	types.LoadTypes(types_file)

	got := getCompiled(t, "../../test/parser_binary_buildin_func.asl")
	want := "someCar setHit [\"motor\", 1];\r\n"

	equal(t, got, want)
}

func TestParserOperator(t *testing.T) {
	got := getCompiled(t, "../../test/parser_operator.asl")
	want := "if (x==y&&x!=y&&x<=y&&x>=y&&x<y&&x>y) then {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserTryCatch(t *testing.T) {
	got := getCompiled(t, "../../test/parser_try_catch.asl")
	want := "try {\r\n} catch {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserNegationFunctionCall(t *testing.T) {
	got := getCompiled(t, "../../test/parser_negation.asl")
	want := "x = !([] call foo);\r\n"

	equal(t, got, want)
}

func TestParserExitWith(t *testing.T) {
	got := getCompiled(t, "../../test/parser_exitwith.asl")
	want := "if (true) exitWith {\r\n};\r\n"

	equal(t, got, want)
}

func TestParserWaitUntil(t *testing.T) {
	got := getCompiled(t, "../../test/parser_waituntil.asl")
	want := "waitUntil {x=x+1;x<100};\r\n"

	equal(t, got, want)
}

func TestParserArray(t *testing.T) {
	got := getCompiled(t, "../../test/parser_array.asl")
	want := "x = [1,2,3];\r\ny = (x select (1));\r\n"

	equal(t, got, want)
}

func TestParserFunctionParams(t *testing.T) {
	got := getCompiled(t, "../../test/parser_func_params.asl")
	want := "myFunc = {\r\nparams [[\"a\",1],[\"b\",2]];\r\nreturn a+b;\r\n};\r\n"

	equal(t, got, want)
}

func TestParserInlineCode(t *testing.T) {
	got := getCompiled(t, "../../test/parser_code.asl")
	want := "inline_code = {a = 1;b = 2;if (a<b) then {[] call foo;};};\r\n"

	equal(t, got, want)
}

func TestParserPreprocessor(t *testing.T) {
	types.LoadTypes(types_file)

	got := getCompiled(t, "../../test/tokenizer_preprocessor.asl")
	want := "\r\n#define HELLO_WORLD \"Hello World!\"\r\nhint HELLO_WORLD;\r\n"

	equal(t, got, want)
}

func TestParserExpressionArray(t *testing.T) {
	got := getCompiled(t, "../../test/parser_expression_array.asl")
	want := "x = [1,2,3]-[2,3];\r\n"

	equal(t, got, want)
}

// bugfix: unary function parsing (e.g. "format")
func TestBugfixParserUnaryFunction(t *testing.T) {
	got := getCompiled(t, "../../test/bugfix_unary_func_format.asl")
	want := "format [\"%1 %2\", \"value1\", \"value2\"];\r\n[\"a\", \"b\", \"c\"] call someFunc;\r\n"

	equal(t, got, want)
}

func getCompiled(t *testing.T, file string) string {
	code, err := ioutil.ReadFile(file)

	if err != nil {
		t.Error("Could not read test file: " + file)
		t.FailNow()
	}

	tokens := tokenizer.Tokenize(code, false)
	compiler := parser.Compiler{}

	return compiler.Parse(tokens, true)
}

func equal(t *testing.T, got, want string) {
	if got != want {
		t.Error("Results do not equal, got:")
		t.Log(got)
		t.Log("expected:")
		t.Log(want)
		t.FailNow()
	}
}
