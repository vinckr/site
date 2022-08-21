package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Page struct {
	File  string
	Title string
	Link  string
}

type PageData struct {
	SiteTitle string
	Title     string
}

var content = [...]Page{
	{
		File: "index",
		Link: "./index.html",
	},
	{
		File: "links",
		Link: "./links.html",
	},
}

func readJson(file string) map[string]interface{} {

	contentpath := "../content/" + file + ".json"
	fmt.Println("Successfully Opened " + contentpath)
	jsonFile, err := ioutil.ReadFile(contentpath)
	check(err)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonFile), &data); err != nil {
		panic(err)
	}
	return data
}

func createPages([2]Page) {
	for k := range content {
		fmt.Println(content[k].File)
		path := "../../static/" + content[k].File + ".html"
		fmt.Println(path)
		template := template.Must(template.New("page").ParseGlob("../templates/*.tmpl"))
		build := new(strings.Builder)

		data := readJson(content[k].File)
		template.Execute(build, data)
		page := build.String()
		println(page)
		os.WriteFile(path, []byte(page), 0666)
	}
}

func main() {
	createPages(content)
}

/*  TODO
    write to several files with same template,
    move content to extra folder
    write to several files with different content
*/
