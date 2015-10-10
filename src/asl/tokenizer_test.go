package asl

import (
    "testing"
    "io/ioutil"
)

func TestTokenizerVar(t *testing.T) {
    got := getTokens(t, "test/tokenizer_var.asl")
	want := []string{"var", "x", "=", "1", ";"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerIf(t *testing.T) {
    got := getTokens(t, "test/tokenizer_if.asl")
	want := []string{"if", "a", "<", "b", "{", "}"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerWhile(t *testing.T) {
    got := getTokens(t, "test/tokenizer_while.asl")
	want := []string{"while", "true", "{", "}"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerFor(t *testing.T) {
    got := getTokens(t, "test/tokenizer_for.asl")
	want := []string{"for", "var", "i", "=", "0", ";", "i", "<", "100", ";", "i", "=", "i", "+", "1", "{", "}"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerEach(t *testing.T) {
    got := getTokens(t, "test/tokenizer_each.asl")
	want := []string{"each", "allUnits", "{", "}"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerFunction(t *testing.T) {
    got := getTokens(t, "test/tokenizer_func.asl")
	want := []string{"func", "TestFunction", "(", "param0", ",", "param1", ")", "{", "return", "true", ";", "}"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerExpression(t *testing.T) {
    got := getTokens(t, "test/tokenizer_expr.asl")
	want := []string{"x", "=", "(", "(", "1", "+", "2", "+", "3", ")", "*", "4", "/", "2", ")", "+", "foo", "(", "1", ",", "2", ",", "3", ")", ";"}
	
	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func compareLength(t *testing.T, got *[]Token, want *[]string) {
    if len(*got) != len(*want) {
	    t.Error("Length of tokens got and expected tokens not equal, was:")
	    gotlist, wantlist := "", ""
	    
	    for i := range *got {
	        gotlist += (*got)[i].token+" "
	    }
	    
	    for i := range *want {
	        wantlist += (*want)[i]+" "
	    }
	    
	    t.Log(gotlist)
	    t.Log("expected:")
	    t.Log(wantlist)
	    t.FailNow()
	}
}

func compareTokens(t *testing.T, got *[]Token, want *[]string) {
    for i := range *got {
	    if (*got)[i].token != (*want)[i] {
	        t.Error("Tokens do not match: "+(*got)[i].token+" != "+(*want)[i])
	    }
	}
}

func getTokens(t *testing.T, file string) []Token {
    code, err := ioutil.ReadFile(file)
    
    if err != nil {
        t.Error("Could not read test file: "+file)
        t.FailNow()
    }
    
	return Tokenize(code)
}
