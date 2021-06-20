package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	var filename string

	flag.StringVar(&filename, "file", "", "Text file name")
	flag.Parse()
	if filename == "" {
		fmt.Println("is empty")
		return
	}

	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(strings.SplitN(filename, ".", 2)[0] + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f, string(fileContents))
	if err != nil {
		panic(err)
	}
	f.Close()
}
