package types

import (
	"io/ioutil"
	"strings"
)

const (
	// type for object types
	TYPE = 1

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
	ArgsCount int
}

var functions []FunctionType

// Returns function type information by name.
// If not found, the first parameter will be false.
func GetFunction(name string) (bool, FunctionType) {
	for _, function := range functions {
		if function.Name == name {
			return true, function
		}
	}

	return false, FunctionType{}
}

// Loads type information from file.
// The format is specified by 'supportInfo' command: https://community.bistudio.com/wiki/supportInfo
func LoadTypes(path string) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	data := strings.Replace(win_new_line, unix_new_line, string(content), -1) // make this work on windows and unix
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

	for _, part := range parts {

	}
}

func parseUnaryFunction(line string) {
	parts := getParts(line)

	for _, part := range parts {

	}
}

func parseBinaryFunction(line string) {
	parts := getParts(line)

	for _, part := range parts {

	}
}

func getParts(line string) []string {
	line = line[2:]
	return strings.Split(line, " ")
}
