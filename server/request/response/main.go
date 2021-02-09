package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Law-key", "this is from lawrence")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(res, "<h1>Any code you want in this func</h1>")

}
func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)

}
