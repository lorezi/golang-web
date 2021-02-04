package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("templates/templates/time-format.gohtml"))
}

func dayMonthYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fdateMDY": dayMonthYear,
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "time-format.gohtml", time.Now())

	if err != nil {
		log.Fatal(err)
	}

}
