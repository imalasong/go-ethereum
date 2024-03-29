package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {

	//printFileNames(".", "\\.+")
	printFileNames("core/vm/runtime", " ├── ", " └── ", ".+_test.go", true, true)
}

func printFileNames(path string, startStr string, endStr string, ignoreReg string, printFile bool, printDir bool) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var dirs []string
	var files []string
	for _, d := range dir {
		if d.IsDir() && printDir {
			if flag, _ := regexp.Match(ignoreReg, []byte(d.Name())); ignoreReg == "" || !flag {
				dirs = append(dirs, d.Name())
			}
		} else if printFile {
			if flag, _ := regexp.Match(ignoreReg, []byte(d.Name())); ignoreReg == "" || !flag {
				files = append(files, d.Name())
			}
		}
	}
	sort.Strings(dirs)
	sort.Strings(files)
	dirs = append(dirs, files...)
	for i, d := range dirs {
		print := d
		if startStr != "" {
			if i < len(dirs)-1 {
				print = startStr + print
			}
		}
		if endStr != "" {
			if i == len(dirs)-1 {
				print = endStr + print
			}
		}
		fmt.Println(print)
	}

}
