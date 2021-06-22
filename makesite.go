package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
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
	var transLang string

	flag.StringVar(&filename, "file", "", "Text file name")
	flag.StringVar(&directory, "dir", "", "File name")
	flag.StringVar(&transLang, "trans", "en", "transLang")
	flag.Parse()

	if directory != "" {
		directoryStuff(directory, transLang)
	} else if filename != "" {
		fileStuff(filename, transLang)
	}

}

func directoryStuff(directory, targetLang string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fileStuff(file.Name(), targetLang)
		}
	}
}

func fileStuff(filename, targetLang string) {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contents := string(fileContents)

	translation, err := translateText(contents, targetLang)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(strings.SplitN(filename, ".", 2)[0] + ".html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(f, translation)
	if err != nil {
		panic(err)
	}
	f.Close()
}

// translateText translates input text and returns translated text.
func translateText(text, targetLang string) (string, error) {

	ctx := context.Background()

	lang, err := language.Parse(targetLang)
	if err != nil {
		return text, fmt.Errorf("targetLangTag: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return text, fmt.Errorf("NewTranslationClient: %v", err)
	}

	translation, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return text, fmt.Errorf("translation: %v", err)
	}

	return translation[0].Text, nil
}
