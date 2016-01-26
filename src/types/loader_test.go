package types_test

import (
	"testing"
	"types"
)

func TestTypesGetFunction(t *testing.T) {
	if err := types.LoadTypes("../../test/types"); err != nil {
		t.Error(err)
	}

	function := types.GetFunction("hint")

	if function == nil {
		t.Error("Function 'hint' not found in type list")
	}
}
