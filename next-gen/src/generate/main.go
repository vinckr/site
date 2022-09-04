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
	Page string
}

var content = [...]Page{
	{
		Page: "index",
	},
	{
		Page: "links",
	},
}

func getContent(Page string) map[string]interface{} {
	path := "../content/" + Page + ".json"
	fmt.Println("Successfully Opened " + path)
	jsonFile, err := ioutil.ReadFile(path)
	check(err)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonFile), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	return data
}

func createPages([2]Page) {
	for k := range content {
		println(content[k].Page)
		path := "../../static/" + content[k].Page + ".html"
		println(path)
		template := template.Must(template.New("page").ParseGlob("../templates/*.tmpl"))
		build := new(strings.Builder)
		data := getContent(content[k].Page)
		var err = template.ExecuteTemplate(build, "page", data)
		check(err)
		if err != nil {
			handleWebErr(build, err)
			return
		}
		page := build.String()
		println(page)
		os.WriteFile(path, []byte(page), 0666)
	}
}

func handleWebErr(build *strings.Builder, err error) {
	panic("unimplemented")
}

func main() {
	createPages(content)
}
