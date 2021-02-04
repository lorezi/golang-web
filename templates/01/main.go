package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Lawrence Onaulogho"

	str := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>` + name + `</h1>
</body>
</html>
	`
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file.", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(str))
}
