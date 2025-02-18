package urlParse

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func GetUrlData(url string) (*html.Node, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("there is some error in get: %w", err)
	}

	defer res.Body.Close()

	body, err := html.Parse(res.Body)

	if err != nil {
		return nil, fmt.Errorf("there is some error reading the response: %w", err)
	}

	return body, nil
}
