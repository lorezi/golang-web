package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/templates/*.gmao"))
}

func main() {
	// Parse requires an absolute path
	// tpl, err := template.ParseGlob("templates/03/templates/*.gmao")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "name.gmao", nil)
	if err != nil {
		log.Fatal(err)
	}

}
