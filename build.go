package main

import (
	"fmt"
	"io/fs"
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

// get files in directory
func getFiles(directory string) []fs.DirEntry {
	files, readDirErr := os.ReadDir(directory)
	if readDirErr != nil {
		log.Fatalf("Error reading files: %s", readDirErr)
	}
	return files
}

// read markdown file
func readMarkdownFile(directory string, filename string) []byte {
	md, readErr := os.ReadFile(directory + filename)
	if readErr != nil {
		log.Fatalf("Error reading file: %s", readErr)
	}
	return md
}

func buildPage(directory string, templates ...string) {

	files := getFiles(directory)

	for _, file := range files {
		md := readMarkdownFile(directory, file.Name())

		// convert markdown to html body
		body := string(markdown.ToHTML(md, nil, nil))
		currentyear := getYear()
		data := data{body, "Blog", "vinckr.com", currentyear, "vinckr"}
		fmt.Print("Pagetitle: " + data.PageTitle)
		// pageData := data{body, data}
		// data := data{body, "Blog", "vinckr.com", currentyear, "vinckr"}

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

	//blogData := data{body, "Blog", "vinckr.com", currentyear, "vinckr"}
	// build blogindex
	//buildPage(blogData, "./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/blog.tmpl")

	// build other pages
	//buildPage("./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

	// build all pages
	buildPage("./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

}
