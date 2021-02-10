package main

import (
	"io"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, "Do my search: "+v)

}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

// visit the page
// https://localhost:8080/?q=dog
