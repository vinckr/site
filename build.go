package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gomarkdown/markdown"
)

// struct for data object
type data struct {
	Body      string
	PageTitle string
	SiteTitle string
	Year      string
	Name      string
}

func getYear() string {
	return time.Now().Format("2006")
}

// declare a ...string
// https://stackoverflow.com/questions/23542989/how-do-i-pass-a-variable-number-of-arguments-to-a-function-in-go

func buildPage(directory string, templates ...string) {

	// get list of files in markdown directory
	files, readDirErr := os.ReadDir(directory)
	if readDirErr != nil {
		log.Fatal("Error reading files: ", readDirErr)
	}

	for _, file := range files {
		// read markdown file
		markdownFile := directory + file.Name()
		content, readErr := os.ReadFile(markdownFile)
		if readErr != nil {
			log.Fatalf("%s file not found", readErr)
		}

		// convert markdown to html body
		body := string(markdown.ToHTML(content, nil, nil))
		currentyear := getYear()
		data := data{body, "Blog", "vinckr.com", currentyear, "vinckr"}
		fmt.Print("Pagetitle: " + data.PageTitle)

		// insert body in template
		var templates = template.Must(template.ParseFiles(templates...))
		build := new(strings.Builder)
		templateErr := templates.ExecuteTemplate(build, "Page", data)
		if templateErr != nil {
			log.Fatalf("Error building the template %s", templateErr)
		}

		// write html file
		outPath := "./public/" + file.Name() + ".html"
		writeErr := ioutil.WriteFile(outPath, []byte(build.String()), 0644)
		if writeErr != nil {
			log.Fatalf("Error writing to %s", writeErr)
		}
		fmt.Printf("\nHTML file %s created \n", outPath)
	}
}

func main() {

	buildPage("./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")
}
