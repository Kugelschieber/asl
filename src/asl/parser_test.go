package asl_test

import (
	"asl"
	"io/ioutil"
	"testing"
)

func TestParserDeclaration(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_var.asl")
	want := "x = 1;\narray = [1,2,3];\n"

	equal(t, got, want)
}

func TestParserAssignment(t *testing.T) {
	got := getCompiled(t, "test/parser_assignment.asl")
	want := "x = 1;\n"

	equal(t, got, want)
}

func TestParserIf(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_if.asl")
	want := "if (a<b) then {\n};\n"

	equal(t, got, want)
}

func TestParserWhile(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_while.asl")
	want := "while {true} do {\n};"

	equal(t, got, want)
}

func TestParserFor(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_for.asl")
	want := "for [{i=0}, {i<100}, {i=i+1}] do {\n};\n"

	equal(t, got, want)
}

func TestParserForeach(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_foreach.asl")
	want := "{\n} forEach (allUnits);\n"

	equal(t, got, want)
}

func TestParserSwitch(t *testing.T) {

}

func TestParserFunction(t *testing.T) {
	got := getCompiled(t, "test/tokenizer_func.asl")
	want := "TestFunction = {\nparam0 = _this select 0;\nparam1 = _this select 1;\nreturn true;\n};\n"

	equal(t, got, want)
}

func TestParserAssignResult(t *testing.T) {
	got := getCompiled(t, "test/parser_assign_result.asl")
	want := "x = ([1, 2, 3] call foo);\ny = ([1, 2, 3] call bar);\n"

	equal(t, got, want)
}

func TestParserExpression(t *testing.T) {
	got := getCompiled(t, "test/parser_expression.asl")
	want := "x = -(1+(2+3))/(6*(someVariable+99-100))-(20)+!anotherVariable+([] call foo);\n"

	equal(t, got, want)
}

func TestParserExpression2(t *testing.T) {
	got := getCompiled(t, "test/parser_expression2.asl")
	want := "x = true||(3>=4&&5<8);\n"

	equal(t, got, want)
}

func TestParserFunctionCall(t *testing.T) {
	got := getCompiled(t, "test/parser_func_call.asl")
	want := "myFunc = {\na = _this select 0;\nb = _this select 1;\nreturn a>b;\n};\n[1+3/4, 2-(66*22)/3-((123))] call myFunc;\n"

	equal(t, got, want)
}

func TestParserBuildinFunctionCall(t *testing.T) {
	got := getCompiled(t, "test/parser_buildin_func.asl")
	want := "_x = (([player, foo] getVar bar) setHit [\"head\", \"tail\"]);\n"

	equal(t, got, want)
}

func TestParserOperator(t *testing.T) {
	got := getCompiled(t, "test/parser_operator.asl")
	want := "if (x==y&&x!=y&&x<=y&&x>=y&&x<y&&x>y) then {\n};\n"

	equal(t, got, want)
}

func TestParserTryCatch(t *testing.T) {
	got := getCompiled(t, "test/parser_try_catch.asl")
	want := "try {\n} catch {\n};\n"

	equal(t, got, want)
}

func TestParserNegationFunctionCall(t *testing.T) {
	got := getCompiled(t, "test/parser_negation.asl")
	want := "x = !([] call foo);\n"

	equal(t, got, want)
}

func TestParserExitWith(t *testing.T) {
	got := getCompiled(t, "test/parser_exitwith.asl")
	want := "if (true) exitWith {\n};\n"

	equal(t, got, want)
}

func TestParserWaitUntil(t *testing.T) {
	got := getCompiled(t, "test/parser_waituntil.asl")
	want := "waitUntil {x=x+1;x<100};\n"

	equal(t, got, want)
}

func TestParserArray(t *testing.T) {
	got := getCompiled(t, "test/parser_array.asl")
	want := "x = [1,2,3];\ny = (x select 1);\n"

	equal(t, got, want)
}

func getCompiled(t *testing.T, file string) string {
	code, err := ioutil.ReadFile(file)

	if err != nil {
		t.Error("Could not read test file: " + file)
		t.FailNow()
	}

	tokens := asl.Tokenize(code)

	return asl.Parse(tokens, true)
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
