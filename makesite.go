package main

import (
	"html/template"
	"os"
	// "fmt"
	// "text/template"
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
	t := template.Must(template.New("template.tmpl").ParseFiles("new.html"))
	err = t.Execute(os.Stdout, todos)
	if err != nil {
		panic(err)
	}
}
