package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome to my home page!")
}

func hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, ps)
	k := ps[0].Key
	v := ps[0].Value
	fmt.Fprintf(w, "key: %s\nValue: %s\n", k, v)
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
