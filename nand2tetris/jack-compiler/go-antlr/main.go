package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"jackcompiler-antlr/compiler"
)

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		abs, _ := filepath.Abs(path)
		fmt.Printf("Could not find file or directory: %s \n", abs)
		os.Exit(1)
	}
	return fileInfo.IsDir()
}

func main() {
	arg := os.Args[1:]

	if len(arg) == 1 {
		path := arg[0]

		if isDirectory(path) {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			for _, f := range files {

				if filepath.Ext(f.Name()) == ".jack" {
					abs, _ := filepath.Abs(path + "/" + f.Name())
					fmt.Printf("Compiling: %s \n", abs)
					compiler.Compile(abs)
				}

			}

		} else {
			abs, _ := filepath.Abs(path)
			fmt.Printf("Compiling: %s \n", abs)
			compiler.Compile(path)
		}

	}
}
