package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Words []string
	Fname string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/global-func.gohtml"))
}

func main() {

	xs := []string{"zero", "one", "two", "three", "four", "five"}

	data := item{
		Words: xs,
		Fname: "Lawrence",
	}

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatal(err)
	}

}
