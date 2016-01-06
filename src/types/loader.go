package types

import (
	"io/ioutil"
	"strings"
)

const (
	// type for object types
	TYPE  = 1
	NAN   = "NaN"
	ARRAY = "ARRAY"

	// types for functions
	NULL   = 2
	UNARY  = 3
	BINARY = 4

	win_new_line  = "\r\n"
	unix_new_line = "\n"
)

type FunctionType struct {
	Name      string
	Type      int // one of the constants NULL, UNARY, BINARY
	ArgsLeft  int
	ArgsRight int // number of args on left side for binary functions
}

var functions []FunctionType

// Returns function type information by name.
// If not found, the parameter will be nil.
func GetFunction(name string) *FunctionType {
	name = strings.ToLower(name)

	for _, function := range functions {
		if function.Name == name {
			return &function
		}
	}

	return nil
}

// Loads type information from file.
// The format is specified by 'supportInfo' command: https://community.bistudio.com/wiki/supportInfo
func LoadTypes(path string) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	data := strings.Replace(win_new_line, unix_new_line, string(content), -1) // make this work on windows and unix
	functions = make([]FunctionType, 0)
	parseTypes(data)

	return nil
}

func parseTypes(content string) {
	lines := strings.Split(content, unix_new_line)

	for _, line := range lines {
		if len(line) < 3 {
			continue
		}

		if line[0] == 'n' {
			parseNullFunction(line)
		} else if line[0] == 'u' {
			parseUnaryFunction(line)
		} else if line[0] == 'b' {
			parseBinaryFunction(line)
		}
	}
}

func parseNullFunction(line string) {
	parts := getParts(line)
	functions = append(functions, FunctionType{parts[0], NULL, 0, 0})
}

func parseUnaryFunction(line string) {
	parts := getParts(line)

	if len(parts) < 2 {
		return
	}

	args := getArgs(parts[1])

	var argsCount int

	if args[0] != ARRAY {
		argsCount = len(args) - getNaNArgs(args)
	}

	functions = append(functions, FunctionType{parts[0], UNARY, argsCount, 0})
}

func parseBinaryFunction(line string) {
	parts := getParts(line)

	if len(parts) < 3 {
		return
	}

	argsLeft := getArgs(parts[0])
	argsRight := getArgs(parts[2])

	var argsLeftCount int
	var argsRightCount int

	if argsLeft[0] != ARRAY {
		argsLeftCount = len(argsLeft) - getNaNArgs(argsLeft)
	}

	if argsRight[0] != ARRAY {
		argsRightCount = len(argsRight) - getNaNArgs(argsRight)
	}

	functions = append(functions, FunctionType{parts[1], BINARY, argsLeftCount, argsRightCount})
}

func getParts(line string) []string {
	line = line[2:]
	return strings.Split(line, " ")
}

func getArgs(part string) []string {
	return strings.Split(part, ",")
}

func getNaNArgs(args []string) int {
	nan := 0

	for _, arg := range args {
		if arg == NAN {
			nan++
		}
	}

	return nan
}
