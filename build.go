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
	PageTitle   string
	Tags        []string
	Log         []string
	Cheatssheet []string
	Other       []string
}

// get files in directory
func getFilesFromDirectory(path string) []fs.DirEntry {
	files, readDirErr := os.ReadDir(path)
	if readDirErr != nil {
		log.Fatalf("Error getting filelist: %s", readDirErr)
	}
	return files
}

// read markdown file from directory
func readMarkdownFileFromDirectory(path string, filename string) []byte {
	md, readErr := os.ReadFile(path + filename)
	if readErr != nil {
		log.Fatalf("Error reading markdown file: %s", readErr)
	}
	return md
}

// split body and frontmatter
func splitBodyAndFrontmatter(md []byte) []byte {
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
func writeHTMLFile(file fs.DirEntry, outpath string, page string) {
	outPath := outpath + strings.TrimSuffix(file.Name(), ".md") + ".html"
	writeErr := ioutil.WriteFile(outPath, []byte(page), 0644)
	if writeErr != nil {
		log.Fatalf("Error writing file: %s", writeErr)
	}
	fmt.Printf("\nHTML file %s created \n", outPath)
}

func buildPages(path string, outpath string, templates ...string) {

	// global config
	author := "vinckr"
	sitetitle := "vinckr.com"
	currentyear := time.Now().Format("2006")

	files := getFilesFromDirectory(path)
	// build pages from files in directory
	for _, file := range files {
		md := readMarkdownFileFromDirectory(path, file.Name())
		bodyOnly := splitBodyAndFrontmatter(md)
		// convert markdown to html body
		body := markdown.ToHTML(bodyOnly, nil, nil)
		// build page object with html body and frontmatter
		page := data{string(body), pagematter.PageTitle, sitetitle, currentyear, author}
		fmt.Printf("Building page %s:", page.PageTitle)
		// build page with template and write to file
		build := buildTemplate(page, templates...)
		writeHTMLFile(file, outpath, build)
	}
}

func main() {

	//build blogindex
	buildPages("./markdown/blog-index/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/blog.tmpl")

	//build home page
	buildPages("./markdown/index/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

	// build all pages
	buildPages("./markdown/content/", "./public/", "./templates/page.tmpl", "./templates/header.tmpl", "./templates/footer.tmpl", "./templates/body.tmpl")

}
