package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/pass-data.gohtml"))
}
func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "pass-data.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatal(err)
	}

}
