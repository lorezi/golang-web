package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/global-func.gohtml"))
}

type score struct {
	Score1 int
	Score2 int
}

func main() {

	g1 := score{
		7, 9,
	}

	err := tpl.Execute(os.Stdout, g1)

	if err != nil {
		log.Fatal(err)
	}

}
