package main

import (
	"io"
	"net/http"
)

type about int

func (m about) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "About Page")
}

type contact int

func (m contact) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Contact Page")

}

type home int

func (m home) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Home Page")
}

func main() {

	var a about
	var c contact
	var h home

	mux := http.NewServeMux()
	// path: root page
	mux.Handle("/", h)
	// path: /about/law/guru
	mux.Handle("/about/", a)
	// path: /contact
	mux.Handle("/contact", c)

	http.ListenAndServe(":8080", mux)

}
