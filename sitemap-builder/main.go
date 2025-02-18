package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html-link-parser/htmlParser"
	"net/url"
	"os"
	"siteMapBuilder/urlParse"
)

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

// URL represents each entry in the sitemap
type URL struct {
	Loc string `xml:"loc"`
}

func getAllLinks(url string) ([]htmlParser.Links, error) {
	body, err := urlParse.GetUrlData(url)

	if err != nil {
		return nil, fmt.Errorf("there is some issue getting data: %w", err)
	}

	document := htmlParser.GetNewHtmlDocumentNode(body)

	return document.GetAllHtmlNodes(), nil
}

func getRootDomain(link string) (string, error) {
	parsedUrl, err := url.Parse(link)

	if err != nil {
		return "", err
	}

	return parsedUrl.Host, nil
}

func bfs(startUrl string) ([]string, error) {
	//keep things in queue once its traversed
	startRootDomain, err := getRootDomain(startUrl)
	if err != nil {
		return nil, fmt.Errorf("there is some issue in parsing the url: %w", err)
	}
	queue := []string{startUrl}
	visited := make(map[string]bool)

	visited[startUrl] = true

	var allLinks []string

	for len(queue) > 0 {
		currentUrl := queue[0]
		queue = queue[1:]

		links, err := getAllLinks(currentUrl)

		if err != nil {
			return nil, fmt.Errorf("there is some issue in getting the data: %w", err)
		}

		for _, link := range links {
			linkDomain, err := getRootDomain(link.Href)
			if err != nil {
				continue
			}

			if linkDomain == startRootDomain && !visited[link.Href] {
				visited[link.Href] = true
				queue = append(queue, link.Href)
				allLinks = append(allLinks, link.Href)
			}
		}
	}

	return allLinks, nil
}

func generateSiteMap(links []string, fileName string) error {
	urlSet := URLSet{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	for _, link := range links {
		urlSet.URLs = append(urlSet.URLs, URL{Loc: link})
	}

	xmlData, err := xml.MarshalIndent(urlSet, "", "  ")
	if err != nil {
		return err
	}

	xmlData = append([]byte(xml.Header), xmlData...)

	err = os.WriteFile(fileName, xmlData, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Sitemap saved to:", fileName)
	return nil
}

func main() {
	var urlParam string
	flag.StringVar(&urlParam, "url", "https://www.calhoun.io", "Enter the url to create the sitemap for")
	flag.Parse()

	links, err := bfs(urlParam)

	if err != nil {
		fmt.Errorf("there is some err: %w", err)
	}

	//Now we have all of the links we will now generate the sitemap for the same
	err = generateSiteMap(append([]string{urlParam}, links...), "sitemap.xml")
	if err != nil {
		fmt.Errorf("there is some error: %w", err)
	}

}
