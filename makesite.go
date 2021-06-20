package main

import (
	// "html/template"
	"io/ioutil"
	"os"

	// "fmt"
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
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	f, err := os.Create("first-post.html")
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
