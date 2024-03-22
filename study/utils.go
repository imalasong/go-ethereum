package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {

	//printFileNames(".", "\\.+")
	printFileNames("cmd", " ├── ", " └── ", "\\.+")
}

func printFileNames(path string, startStr string, endStr string, ignoreReg string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var ss []string
	for _, d := range dir {
		if d.IsDir() {
			if flag, _ := regexp.Match(ignoreReg, []byte(d.Name())); ignoreReg == "" || !flag {
				ss = append(ss, d.Name())
			}
		}
	}
	sort.Strings(ss)
	for i, d := range ss {
		print := d
		if startStr != "" {
			if i < len(ss)-1 {
				print = startStr + print
			}
		}
		if endStr != "" {
			if i == len(ss)-1 {
				print = endStr + print
			}
		}
		fmt.Println(print)
	}

}
