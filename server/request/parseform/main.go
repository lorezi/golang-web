package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

type data struct {
	Method        string
	Submissions   url.Values
	URL           *url.URL
	Header        http.Header
	Host          string
	ContentLength int64
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	d := data{
		r.Method,
		r.Form,
		r.URL,
		r.Header,
		r.Host,
		r.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", d)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("server/request/parseform/index.gohtml"))
}

func main() {
	var s hotdog
	http.ListenAndServe(":8080", s)
}
