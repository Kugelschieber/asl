package tokenizer_test

import (
    "tokenizer"
	"io/ioutil"
	"testing"
)

func TestTokenizerVar(t *testing.T) {
	got := getTokens(t, "test/tokenizer_var.asl")
	want := []string{"var", "x", "=", "1", ";", "var", "array", "=", "[", "1", ",", "2", ",", "3", "]", ";"}

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

func TestTokenizerForach(t *testing.T) {
	got := getTokens(t, "test/tokenizer_foreach.asl")
	want := []string{"foreach", "allUnits", "{", "}"}

	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerSwitch(t *testing.T) {
	got := getTokens(t, "test/tokenizer_switch.asl")
	want := []string{"switch", "x", "{", "case", "1", ":", "x", "=", "1", ";", "case", "2", ":", "x", "=", "2", ";", "default", ":", "x", "=", "3", ";", "}"}

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

func TestTokenizerIdentifier(t *testing.T) {
	got := getTokens(t, "test/tokenizer_identifier.asl")
	want := []string{"var", "format", "=", "\"should not be for mat!\"", ";"}

	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func TestTokenizerInlineCode(t *testing.T) {
	got := getTokens(t, "test/tokenizer_code.asl")
	want := []string{"var", "x", "=", "code", "(", "\"var x = 5;\"", ")", ";"}

	compareLength(t, &got, &want)
	compareTokens(t, &got, &want)
}

func compareLength(t *testing.T, got *[]tokenizer.Token, want *[]string) {
	if len(*got) != len(*want) {
		t.Error("Length of tokens got and expected tokens not equal, was:")
		gotlist, wantlist := "", ""

		for i := range *got {
			gotlist += (*got)[i].Token + " "
		}

		for i := range *want {
			wantlist += (*want)[i] + " "
		}

		t.Log(gotlist)
		t.Log("expected:")
		t.Log(wantlist)
		t.FailNow()
	}
}

func compareTokens(t *testing.T, got *[]tokenizer.Token, want *[]string) {
	for i := range *got {
		if (*got)[i].Token != (*want)[i] {
			t.Error("Tokens do not match: " + (*got)[i].Token + " != " + (*want)[i])
		}
	}
}

func getTokens(t *testing.T, file string) []tokenizer.Token {
	code, err := ioutil.ReadFile(file)

	if err != nil {
		t.Error("Could not read test file: " + file)
		t.FailNow()
	}

	return tokenizer.Tokenize(code)
}
