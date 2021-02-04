package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Name  string
	Model string
}

type items struct {
	Wisdom    []sage
	Transport []car
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/templates/funcTemp.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}

	return s
}

func main() {

	b := sage{
		Name:  "Bu",
		Motto: "The belief of no belief.",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change.",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred.",
	}

	t := car{
		Name:  "Tesla",
		Model: "G459",
	}

	f := car{
		Name:  "Ferrari",
		Model: "GT098",
	}

	sages := []sage{b, g, m}
	cars := []car{t, f}

	data := items{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "funcTemp.gohtml", data)

	if err != nil {
		log.Fatal(err)
	}

}
