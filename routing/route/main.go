package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/about":
		io.WriteString(w, "About me page")
	case "/contact":
		io.WriteString(w, "Contact page")
	default:
		io.WriteString(w, "Home page")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
