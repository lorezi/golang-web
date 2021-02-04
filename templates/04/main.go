package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/data.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "data.gohtml", 45)
	if err != nil {
		log.Fatal(err)
	}
}
