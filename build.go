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

	"github.com/adrg/frontmatter"
	"github.com/gomarkdown/markdown"
)

// struct for data object including content and global config
type data struct {
	Body      string
	PageTitle string
	SiteTitle string
	Year      string
	Author    string
}

// struct for frontmatter
var pagematter struct {
	PageTitle string
	Tags      []string
	Linklist  []string
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

// split body and frontmatter
func bodyOnly(md []byte) []byte {
	bodyOnly, err := frontmatter.Parse(strings.NewReader(string(md)), &pagematter)
	if err != nil {
		log.Fatalf("Error parsing frontmatter: %s", err)
	}
	return bodyOnly
}

// insert body in template
func buildTemplate(data data, templates ...string) string {
	var t = template.Must(template.ParseFiles(templates...))
	build := new(strings.Builder)
	templateErr := t.ExecuteTemplate(build, "Page", data)
	if templateErr != nil {
		log.Fatalf("Error building the template %s", templateErr)
	}
	return build.String()
}

// write html file
func writeHTMLFile(file fs.DirEntry, page string) {
	outPath := "./public/" + strings.TrimSuffix(file.Name(), ".md") + ".html"
	writeErr := ioutil.WriteFile(outPath, []byte(page), 0644)
	if writeErr != nil {
		log.Fatalf("Error writing file: %s", writeErr)
	}
	fmt.Printf("\nHTML file %s created \n", outPath)
}

func buildPage(directory string, templates ...string) {

	// global config
	author := "vinckr"
	sitetitle := "vinckr.com"
	currentyear := time.Now().Format("2006")

	files := getFiles(directory)
	// build pages from files in directory
	for _, file := range files {
		md := readMarkdownFile(directory, file.Name())
		bodyOnly := bodyOnly(md)
		// convert markdown to html body
		body := markdown.ToHTML(bodyOnly, nil, nil)
		// build page object with html body and frontmatter
		page := data{string(body), pagematter.PageTitle, sitetitle, currentyear, author}
		// build page with template and write to file
		build := buildTemplate(page, templates...)
		writeHTMLFile(file, build)
	}
}

func main() {

	// build blogindex
	//buildPage(blogData, "./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/blog.tmpl")

	// build other pages
	//buildPage("./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

	// build all pages
	buildPage("./markdown/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

}
