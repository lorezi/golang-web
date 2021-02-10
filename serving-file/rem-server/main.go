package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")

	io.WriteString(w, `<img src="https://picsum.photos/200/300"/>`)
}

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)

}
