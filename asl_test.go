package main

import (
	"io/ioutil"
	"testing"
)

type Want struct {
    tokens []string
    parser string
}

type Got struct {
    tokens []Token
    parser string
}

func TestArray(t *testing.T) {
	got := getCompiled(t, "test/array.asl")
    want := &Want{
        []string{"var","x","=","[","1",",","2",",","3","]",";","var","y","=","x","[","1","]",";"},
        "x = [1,2,3];\r\ny = (x select (1));\r\n",
    }

	equal(t, got, want)
}

func TestAssignResult(t *testing.T) {
	got := getCompiled(t, "test/assign_result.asl")
    want := &Want{
        []string{"var", "x", "=", "foo", "(", "1", ",", "2", ",", "3", ")", ";", "y", "=", "bar", "(", "1", ",", "2", ",", "3", ")", ";"},
	    "x = ([1, 2, 3] call foo);\r\ny = ([1, 2, 3] call bar);\r\n",
    }

	equal(t, got, want)
}

func TestAssignment(t *testing.T) {
	got := getCompiled(t, "test/assignment.asl")
    want := &Want{
        []string{"x", "=", "1", ";"},
	    "x = 1;\r\n",
    }

	equal(t, got, want)
}

func TestBuildinFunctionCall(t *testing.T) {
	got := getCompiled(t, "test/buildin_func.asl")
    want := &Want{
        []string{"var","_x","=","setHit","(","getVar","(","player",",","foo",")","(","bar",")",")","(","\"head\"",",","\"tail\"",")",";"},
	    "_x = (([player, foo] getVar bar) setHit [\"head\", \"tail\"]);\r\n",
    }

	equal(t, got, want)
}

func TestCode(t *testing.T) {
	got := getCompiled(t, "test/code.asl")
    want := &Want{
        []string{"var", "x", "=", "code", "(", "\"var x = 5;\"", ")", ";"},
	    "x = {x = 5;};\r\n",
    }

	equal(t, got, want)
}

func TestComment(t *testing.T) {
	got := getCompiled(t, "test/comment.asl")
    want := &Want{
        []string{"var","x","=","1",";"},
	    "x = 1;\r\n",
    }

	equal(t, got, want)
}


func TestExitWith(t *testing.T) {
	got := getCompiled(t, "test/exitwith.asl")
    want := &Want{
        []string{"exitwith","{","}"},
	    "if (true) exitWith {\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestExpression1(t *testing.T) {
	got := getCompiled(t, "test/expression1.asl")
    want := &Want{
        []string{"x","=","(","(","1","+","2","+","3",")","*","4","/","2",")","+","foo","(","1",",","2",",","3",")",";"},
	    "x = ((1+2+3)*4/2)+([1, 2, 3] call foo);\r\n",
	}
	equal(t, got, want)
}

func TestExpression2(t *testing.T) {
	got := getCompiled(t, "test/expression2.asl")
    want := &Want{
        []string{"var","x","=","true","|","|","(","3",">","=","4","&","&","5","<","8",")",";"},
	    "x = true||(3>=4&&5<8);\r\n",
    }

	equal(t, got, want)
}

func TestExpression3(t *testing.T) {
	got := getCompiled(t, "test/expression3.asl")
    want := &Want{
        []string{"var","x","=","-","(","1","+","(","2","+","3",")",")","/","(","6","*","(","someVariable","+","99","-","100",")",")","-","(","20",")","+","!","anotherVariable","+","foo","(",")",";"},
	    "x = -(1+(2+3))/(6*(someVariable+99-100))-(20)+!anotherVariable+([] call foo);\r\n",
    }

	equal(t, got, want)
}

func TestFor(t *testing.T) {
	got := getCompiled(t, "test/for.asl")
    want := &Want{
        []string{"for","var","i","=","0",";","i","<","100",";","i","=","i","+","1","{","}"},
	    "for [{i=0}, {i<100}, {i=i+1}] do {\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestForeach(t *testing.T) {
	got := getCompiled(t, "test/foreach.asl")
    want := &Want{
        []string{"foreach", "unit", "=", ">", "allUnits", "{", "}"},
	    "{\r\nunit = _x;\r\n} forEach (allUnits);\r\n",
    }

	equal(t, got, want)
}

func TestFunction(t *testing.T) {
	got := getCompiled(t, "test/func.asl")
    want := &Want{
        []string{"func", "TestFunction", "(", "param0", ",", "param1", ")", "{", "return", "true", ";", "}"},
	    "TestFunction = {\r\nparams [\"param0\",\"param1\"];\r\nreturn true;\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestFunctionCall(t *testing.T) {
	got := getCompiled(t, "test/func_call.asl")
    want := &Want{
        []string{"func","myFunc","(","a",",","b",")","{","return","a",">","b",";","}","myFunc","(","1","+","3","/","4",",","2","-","(","66","*","22",")","/","3","-","(","(","123",")",")",")",";"},
	    "myFunc = {\r\nparams [\"a\",\"b\"];\r\nreturn a>b;\r\n};\r\n[1+3/4, 2-(66*22)/3-((123))] call myFunc;\r\n",
    }

	equal(t, got, want)
}

func TestFunctionParams(t *testing.T) {
	got := getCompiled(t, "test/func_params.asl")
    want := &Want{
        []string{"func","myFunc","(","a","=","1",",","b","=","2",")","{","return","a","+","b",";","}"},
	    "myFunc = {\r\nparams [[\"a\",1],[\"b\",2]];\r\nreturn a+b;\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestIdentifier(t *testing.T) {
	got := getCompiled(t, "test/identifier.asl")
    want := &Want{
        []string{"var","format","=","\"should not be for mat!\"",";"},
	    "format = \"should not be for mat!\";\r\n",
    }

	equal(t, got, want)
}

func TestIf(t *testing.T) {
	got := getCompiled(t, "test/if.asl")
    want := &Want{
        []string{"if","a","<","b","{","}"},
	    "if (a<b) then {\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestInlineCode(t *testing.T) {
	got := getCompiled(t, "test/inline_code.asl")
    want := &Want{
        []string{"var","inline_code","=","code","(","\"var a = 1;var b = 2;if a < b {foo();}\"",")",";"},
	    "inline_code = {a = 1;b = 2;if (a<b) then {[] call foo;};};\r\n",
    }

	equal(t, got, want)
}

func TestNegation(t *testing.T) {
	got := getCompiled(t, "test/negation.asl")
    want := &Want{
        []string{"var","x","=","!","foo","(",")",";"},
	    "x = !([] call foo);\r\n",
	}

	equal(t, got, want)
}

func TestOperator(t *testing.T) {
	got := getCompiled(t, "test/operator.asl")
    want := &Want{
        []string{"if","x","=","=","y","&","&","x","!","=","y","&","&","x","<","=","y","&","&","x",">","=","y","&","&","x","<","y","&","&","x",">","y","{","}"},
	    "if (x==y&&x!=y&&x<=y&&x>=y&&x<y&&x>y) then {\r\n};\r\n",
    }


	equal(t, got, want)
}

func TestPreprocessor(t *testing.T) {
	got := getCompiled(t, "test/preprocessor.asl")
    want := &Want{
        []string{"#define HELLO_WORLD \"Hello World!\"", "hint", "(", ")", "(", "HELLO_WORLD", ")", ";"},
	    "\r\n#define HELLO_WORLD \"Hello World!\"\r\nhint HELLO_WORLD;\r\n",
    }

	equal(t, got, want)
}

func TestSwitch(t *testing.T) {
    got := getCompiled(t, "test/switch.asl")
    want := &Want{
        []string{"switch","x","{","case","1",":","x","=","1",";","case","2",":","x","=","2",";","default",":","x","=","3",";","}"},
	    "switch (x) do {\r\ncase 1:\r\n{\r\nx = 1;\r\n};\r\ncase 2:\r\n{\r\nx = 2;\r\n};\r\ndefault:\r\n{\r\nx = 3;\r\n};\r\n};\r\n",
    }

	equal(t, got, want)
}

func TestTryCatch(t *testing.T) {
	got := getCompiled(t, "test/try_catch.asl")
    want := &Want{
        []string{"try","{","}","catch","{","}"},
	    "try {\r\n} catch {\r\n};\r\n",
    }
	equal(t, got, want)
}

func TestWaitUntil(t *testing.T) {
	got := getCompiled(t, "test/waituntil.asl")
    want := &Want{
        []string{"waituntil","(","x","=","x","+","1",";","x","<","100",")",";"},
	    "waitUntil {x=x+1;x<100};\r\n",
    }

	equal(t, got, want)
}


func TestWhile(t *testing.T) {
	got := getCompiled(t, "test/while.asl")
    want := &Want{
        []string{"while", "true", "{", "}"},
	    "while {true} do {\r\n};",
    }

	equal(t, got, want)
}

func getCompiled(t *testing.T, file string) *Got {
	code, err := ioutil.ReadFile(file)

	if err != nil {
		t.Error("Could not read test file: " + file)
		t.FailNow()
	}

	tokens := Tokenize(code)
	compiler := Compiler{}
	parsed := compiler.Parse(tokens, true)

	got := &Got{tokens, parsed}

	return got
}

func compareLength(t *testing.T, got *Got, want *Want) {
	if len(got.tokens) != len(want.tokens) {
		t.Error("Length of tokens got and expected tokens not equal, was:")
		gotlist, wantlist := "", ""

		for i := range got.tokens {
			gotlist += (got.tokens)[i].Token + " "
		}

		for i := range want.tokens {
			wantlist += (want.tokens)[i] + " "
		}

		t.Log(gotlist)
		t.Log("expected:")
		t.Log(wantlist)
		t.FailNow()
	}


}

func compareTokens(t *testing.T, got *Got, want *Want) {
	for i := range got.tokens {
		if (got.tokens)[i].Token != (want.tokens)[i] {
			t.Error("Tokens do not match: " + (got.tokens)[i].Token + " != " + (want.tokens)[i])
		}
	}
}

func equal(t *testing.T, got *Got, want *Want) {
	compareLength(t, got, want)
	compareTokens(t, got, want)

	if got.parser != want.parser {
		t.Error("Parsed does not equal, got:")
		t.Log(got.parser)
		t.Log("expected:")
		t.Log(want.parser)
		t.FailNow()
	}
}
