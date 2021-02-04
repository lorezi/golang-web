package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/composition.gohtml"))
}

func main() {
	p1 := person{
		Name: "John Doe",
		Age:  44,
	}

	err := tpl.Execute(os.Stdout, p1)

	if err != nil {
		log.Fatal(err)
	}

}
