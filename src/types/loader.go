package types

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Loads type information from file.
func LoadTypes(path string) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	if err := parseTypes(string(content)); err != nil {
		return err
	}

	return nil
}

func parseTypes(content string) error {
	lines := strings.Split(content, "\\n")

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
