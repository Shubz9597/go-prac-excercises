package main

import (
	"fmt"
	"html-link-parser/htmlParser"
	"log"
)

func main() {
	parsed, err := htmlParser.ParseHtml("ex2")
	if err != nil {
		log.Fatal(err)
	}

	document := htmlParser.GetNewHtmlDocumentNode(parsed)

	tags := document.GetAllHtmlNodes()

	fmt.Print(tags)
}
