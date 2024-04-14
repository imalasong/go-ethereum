package main

import (
	bytes2 "bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"text/template"
)

func main() {

	temp := `Name :{{.Name}}
Age :{{.Age | toYear}}
gender:{{.Gender}}`

	funcs := map[string]interface{}{
		"toYear": toYear,
	}
	parse, err := template.New("txt").Funcs(funcs).Parse(temp)

	if err != nil {
		log.Error("parse template error:", err)
	}
	p := Person2{"aa", 1, 1}

	//err = parse.Execute(os.Stdout, p)
	bytes := new(bytes2.Buffer)
	err = parse.Execute(bytes, p)

	if err != nil {
		log.Error("execute template error:", err)
	}
	fmt.Printf("%s", bytes.Bytes())
}

type Person2 struct {
	Name   string `json:"name"`
	Age    int
	Gender int
}

func toYear(age int) int {
	return age * 100
}
