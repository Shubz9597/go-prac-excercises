package htmlParser

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type HtmlDocument struct {
	Root *html.Node
}

type Links struct {
	Href string
	Text string
}

func ParseHtml(fileName string) (*html.Node, error) {

	file := "html/" + fileName + ".html"
	reader, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("there is some error parsing the file %w", err)
	}

	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("there is some error parsing the html file %w", err)
	}

	return doc, nil
}

func GetNewHtmlDocumentNode(node *html.Node) *HtmlDocument {
	return &HtmlDocument{Root: node}
}

func (d *HtmlDocument) GetAllHtmlNodes() []Links {

	var tags []Links
	var traverseTree func(n *html.Node)

	traverseTree = func(n *html.Node) {
		if n.Data == "a" && n.Type == html.ElementNode {
			href := ""

			for _, links := range n.Attr {
				if links.Key == "href" {
					href = links.Val
					break
				}
			}

			text := strings.TrimSpace(extractString(n))
			tags = append(tags, Links{Href: href, Text: text})
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseTree(c)
		}
	}

	traverseTree(d.Root)
	return tags
}

func extractString(n *html.Node) string {
	var str string = ""

	for i := n.FirstChild; i != nil; i = i.NextSibling {

		if i.Data == "a" && i.Type == html.ElementNode {
			continue
		}

		if i.Type == html.TextNode {
			str += i.Data
		}

		if i.Type == html.ElementNode {
			str += extractString(i)
		}
	}

	return str
}
