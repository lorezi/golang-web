package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/slice.gohtml"))
}
func main() {

	sages := []string{"Gandhi", "MLK", "Buddha", "Muhammad"}
	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatal(err)
	}

}
