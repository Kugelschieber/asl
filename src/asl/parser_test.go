package asl

import (
    "testing"
    "io/ioutil"
)

func TestParserDeclaration(t *testing.T) {
    got := getCompiled(t, "test/tokenizer_var.asl")
    want := "x = 1;\n"
    
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

func TestParserEach(t *testing.T) {
    got := getCompiled(t, "test/tokenizer_each.asl")
    want := "{\n} forEach (allUnits);\n"
    
    equal(t, got, want)
}

func TestParserFunction(t *testing.T) {
    got := getCompiled(t, "test/tokenizer_func.asl")
    want := "TestFunction = {\nparam0 = _this select 0;\nparam1 = _this select 1;\nreturn true;\n};\n"
    
    equal(t, got, want)
}

// TODO
/*func TestParserAssignResult(t *testing.T) {
    got := getCompiled(t, "test/parser_assign_result.asl")
    want := "x = [1, 2, 3] call foo;\ny = [1, 2, 3] call bar;"
    
    equal(t, got, want)
}*/

func getCompiled(t *testing.T, file string) string {
    code, err := ioutil.ReadFile(file)
    
    if err != nil {
        t.Error("Could not read test file: "+file)
        t.FailNow()
    }
    
	tokens := Tokenize(code)
	
	return Parse(tokens, true)
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
