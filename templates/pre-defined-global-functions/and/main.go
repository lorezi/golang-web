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

type user struct {
	Name  string
	Role  string
	Admin bool
}

func main() {

	u1 := user{
		Name:  "Fejiro Tega",
		Role:  "System Developer",
		Admin: true,
	}

	u2 := user{
		Name:  "Haruna Victor",
		Role:  "HR Admin",
		Admin: false,
	}

	u3 := user{
		Name:  "",
		Role:  "Nil",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)

	if err != nil {
		log.Fatal(err)
	}

}
