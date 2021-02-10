package main

import (
	"io"
	"net/http"
)

// passing values from form
// method post sends the values through the body
// method get sends the values through the url

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
		<input type="text" name="q"/>
		<input type="submit"/>
	</form>
	<br/>
	`+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
