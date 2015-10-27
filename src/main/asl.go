package main

import (
	"asl"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"os"
	"strings"
)

const version = "1.0.0"
const extension = ".asl"
const sqfextension = ".sqf"

type ASLFile struct {
    in string
    out string
    newname string
}

var recursive bool = false
var pretty bool = false
var exit bool = false
var function bool = false;
var aslFiles []ASLFile
var inDir string

func usage() {
	fmt.Println("Usage: asl [-v|-r|-pretty|--help] <input directory> <output directory>\n")
	fmt.Println("-v (optional) shows asl version")
	fmt.Println("-r (optional) recursivly compile all asl files in folder")
	fmt.Println("-pretty (optional) activates pretty printing\n")
	fmt.Println("--help (optional) shows usage\n")
	fmt.Println("<input directory> directory to compile")
	fmt.Println("<output directory> output directory, directory structure will be created corresponding to input directory")
}

func flags(flag string) bool {
    flag = strings.ToLower(flag)
    
    if flag[0] == '-' {
        if flag == "-v" {
            fmt.Println("asl version "+version)
            exit = true
        } else if flag == "-r" {
            recursive = true
        } else if flag == "-f" {
            function = true
        } else if flag == "-pretty" {
            pretty = true
        } else if flag == "--help" {
            usage()
            exit = true
        }
        
        return true
    }
    
    return false
}

func readAslFiles(path string) {
    dir, err := ioutil.ReadDir(path)
    
    if err != nil {
        fmt.Println("Error reading in directory!")
        return
    }
    
    for i := 0; i < len(dir); i++ {
        name := dir[i].Name()
        
        if dir[i].IsDir() && recursive {
            readAslFiles(filepath.FromSlash(path+"/"+name))
            continue
        }

        if !dir[i].IsDir() && strings.ToLower(filepath.Ext(name)) == extension {
            in := filepath.FromSlash(path+"/"+dir[i].Name())
            newname := name[:len(name)-len(filepath.Ext(name))]

            out := filepath.FromSlash("./"+path[len(inDir):len(path)])
            if function {
                out = filepath.FromSlash("./functions/"+path[len(inDir):len(path)])
            }   
            
            file := ASLFile{in, out, newname}
            aslFiles = append(aslFiles, file)
        }
    }
}

func compile(path string) {
    for i := 0; i < len(aslFiles); i++ {
        out := filepath.FromSlash(path+aslFiles[i].out+aslFiles[i].newname+sqfextension)
        if function {
            out = filepath.FromSlash(path+aslFiles[i].out+"fn_"+aslFiles[i].newname+sqfextension)
        }
        
        fmt.Println(aslFiles[i].in+" -> "+out)
        code, err := ioutil.ReadFile(aslFiles[i].in)
        if err != nil {
            fmt.Println("Error reading file: "+aslFiles[i].in)
            continue
        }
        
    	token := asl.Tokenize(code)
    	sqf := asl.Parse(token, pretty)
    	
    	os.MkdirAll(filepath.FromSlash(path+"/"+aslFiles[i].out), 0777)
    	err = ioutil.WriteFile(out, []byte(sqf), 0666)
    	
    	if err != nil {
    	    fmt.Println("Error writing file: "+aslFiles[i].out)
    	    fmt.Println(err)
    	}
    }    
}

func functions(path string) {
    functionText := "//class cfgFunctions {\nclass asl {\n    class asl {\n"

    for i := 0; i < len(aslFiles); i++ {
        functionText +="        class "+aslFiles[i].newname+";\n"
    }

    functionText +="    }\n}\n//}"
    ioutil.WriteFile(filepath.FromSlash(path+"/"+"functions.hpp"), []byte(functionText), 0666)
    fmt.Println("functions"+" -> "+path+"/"+"functions.hpp")
}

func main() {
	args := os.Args
	
	// flags
	if len(args) < 2 {
	    usage()
	    return
	}
	
	var i int
	for i = 1; i < len(args) && flags(args[i]); i++ {}
	
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
        i++
	}
	
	readAslFiles(inDir)

    if len(aslFiles) > 0 {
        compile(out)
    } else {
        fmt.Println("No asl files found from "+inDir)
        return
    }
	
    if function {
        functions(out)
    }
}
