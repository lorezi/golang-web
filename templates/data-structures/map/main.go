package main

import (
	"log"
	"os"
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

// struct of a slice of struct
/*
{
	struct
	[ slice
		{
			struct
		}
	]
}

*/
type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.ParseFiles("templates/templates/map.gohtml"))
}
func main() {
	// sages := map[string]string{
	// 	"India":    "Gandhi",
	// 	"America":  "MLK",
	// 	"Meditate": "Buddha",
	// 	"Prophet":  "Muhammad",
	// }

	// struct
	// buddha := sage{
	// 	Name:  "Buddha",
	// 	Motto: "The belief of no beliefs",
	// }

	// slice of struct
	// buddha := sage{
	// 	Name:  "Buddha",
	// 	Motto: "The belief of no belief",
	// }

	// gandhi := sage{
	// 	Name:  "Gandhi",
	// 	Motto: "Be the change",
	// }

	// mlk := sage{
	// 	Name:  "Martin Luther King",
	// 	Motto: "Hatred never ceases with hatred",
	// }
	// // err := tpl.Execute(os.Stdout, buddha)
	// sages := []sage{buddha, gandhi, mlk}

	// err := tpl.Execute(os.Stdout, sages)

	// struct-slice-struct

	b := sage{
		Name:  "Buddha",
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

	err := tpl.Execute(os.Stdout, data)

	if err != nil {
		log.Fatal(err)
	}

}
