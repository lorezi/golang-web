package main

import (
	"fmt"
	"net/http"
)

func h(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Welcome to my home page")
}

func a(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "About page of this exercise.")
}

func c(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Contact me for golang jobs 😀")
}

func main() {
	http.HandleFunc("/", h)
	http.HandleFunc("/about/:name", a)
	http.HandleFunc("/contact", c)

	http.ListenAndServe(":8080", nil)

}
