// final version

package main

import (
	"io"
	"net/http"
)

func h(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Home Page")
}

func a(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "About Page")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Contact Page")
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/about/", a)
	http.HandleFunc("/contact", c)

	http.ListenAndServe(":8080", nil)

}
