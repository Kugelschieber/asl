package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"parser"
	"path/filepath"
	"strings"
	"tokenizer"
	"types"
)

const (
	version       = "1.2.2"
	extension     = ".asl"
	sqfextension  = ".sqf"
	typeinfo      = "types"
	PathSeparator = string(os.PathSeparator)
)

type ASLFile struct {
	in      string
	out     string
	newname string
}

var (
	recursive bool = false
	pretty    bool = false
	exit      bool = false
	aslFiles  []ASLFile
	inDir     string
)

func usage() {
	fmt.Println("Usage: asl [-v|-r|-pretty|--help] <input directory> <output directory>\n")
	fmt.Println("-v (optional) shows asl version")
	fmt.Println("-r (optional) recursivly compile all asl files in folder")
	fmt.Println("-pretty (optional) activates pretty printing\n")
	fmt.Println("--help (optional) shows usage\n")
	fmt.Println("<input directory> directory to compile")
	fmt.Println("<output directory> output directory, directory structure will be created corresponding to input directory")
}

// Parses compiler flags.
func flags(flag string) bool {
	flag = strings.ToLower(flag)

	if flag[0] != '-' {
		return false
	}

	if flag == "-v" {
		fmt.Println("asl version " + version)
		exit = true
	} else if flag == "-r" {
		recursive = true
	} else if flag == "-pretty" {
		pretty = true
	} else if flag == "--help" {
		usage()
		exit = true
	}

	return true
}

// Loads types from types file.
// If none is provided, an error will be printed.
func loadTypes() {
	if err := types.LoadTypes(typeinfo); err != nil {
		fmt.Println("No 'types' file provided. Please add type information to this file from 'supportInfo' script command output.")
		exit = true
	}
}

// Creates a list of all ASL files to compile.
func readAslFiles(path string) {
	dir, err := ioutil.ReadDir(path)

	if err != nil {
		fmt.Println("Error reading in directory!")
		return
	}

	for i := 0; i < len(dir); i++ {
		name := dir[i].Name()

		if dir[i].IsDir() && recursive {
			readAslFiles(filepath.FromSlash(path + PathSeparator + name))
			continue
		}

		if !dir[i].IsDir() && strings.ToLower(filepath.Ext(name)) == extension {
			in := filepath.FromSlash(path + PathSeparator + dir[i].Name())
			out := filepath.FromSlash("./" + path[len(inDir):len(path)])
			newname := name[:len(name)-len(filepath.Ext(name))]

			file := ASLFile{in, out, newname}
			aslFiles = append(aslFiles, file)
		}
	}
}

// Recovers and prints thrown error.
func recoverCompileError(file string) {
	if r := recover(); r != nil {
		fmt.Println("Compile error in file "+file+":", r)
	}
}

// Compiles a single ASL file.
func compileFile(path string, file ASLFile) {
	defer recoverCompileError(file.in)

	// read file
	out := filepath.FromSlash(path + PathSeparator + file.out + PathSeparator + file.newname + sqfextension)
	fmt.Println(file.in + " -> " + out)
	code, err := ioutil.ReadFile(file.in)

	if err != nil {
		fmt.Println("Error reading file: " + file.in)
		return
	}

	// compile
	token := tokenizer.Tokenize(code, false)
	compiler := parser.Compiler{}
	sqf := compiler.Parse(token, pretty)

	os.MkdirAll(filepath.FromSlash(path+PathSeparator+file.out), 0777)
	err = ioutil.WriteFile(out, []byte(sqf), 0666)

	if err != nil {
		fmt.Println("Error writing file: " + file.out)
		fmt.Println(err)
	}
}

// Compiles ASL files.
func compile(path string) {
	for i := 0; i < len(aslFiles); i++ {
		compileFile(path, aslFiles[i])
	}
}

func main() {
	args := os.Args

	// flags
	if len(args) < 2 {
		usage()
		return
	}

	var i int
	for i = 1; i < len(args) && flags(args[i]); i++ {
	}

	if exit {
		return
	}

	// load type information
	loadTypes()

	if exit {
		return
	}

	// in/out parameter
	out := ""

	if i < len(args) {
		inDir = args[i]
		i++
	} else {
		return
	}

	if i < len(args) {
		out = args[i]
	}

	readAslFiles(inDir)
	compile(out)
}
