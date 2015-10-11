package main

import (
	"fmt"
	"path/filepath"
)

type ValidateFile func(string) (string, bool)

type Info struct {
	Vfile ValidateFile
}

func checkFile(file string) (string, bool) {
	pathFile, _ := filepath.Abs(file)
	checked := false
	if pathFile != "" {
		return pathFile, true
	}
	return "", checked
}

func main() {

	i := Info{}
	i.Vfile = checkFile

	if path, ok := i.Vfile("testing"); ok {
		fmt.Println(path)
	}
}
