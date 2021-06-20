package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	// "log"
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
	var directory string

	flag.StringVar(&filename, "file", "", "Text file name")
	flag.StringVar(&directory, "dir", "", "File name")
	flag.Parse()

	if directory != "" {
		directoryStuff(directory)
	} else if filename != "" {
		fileStuff(filename)
	}

}

func directoryStuff(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fileStuff(file.Name())
		}
	}
}

func fileStuff(filename string) {
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
