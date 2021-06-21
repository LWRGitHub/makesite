package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
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
	flag.StringVar(&directory, "trans", "", "File name")
	flag.Parse()

	if directory != "" {
		directoryStuff(directory)
	} else if filename != "" {
		fileStuff(filename)
	} else if transLang != "" {
		trans1(filename)
	}

}

func trans1(filename string) {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fileContents = translateText(fileContents)

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

// translateText translates input text and returns translated text.
func translateText(w io.Writer, text string) error {
	// w := io.Writer
	projectID := "carbide-program-317404"
	sourceLang := "en-US"
	targetLang := "fr"
	// text := "Text you wish to translate"

	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return fmt.Errorf("TranslateText: %v", err)
	}

	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
		fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
	}

	return nil
}
