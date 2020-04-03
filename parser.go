package linkparser

import (
	"io"
	"golang.org/x/net/html"
	"fmt"
)

// Link represents an HTML link (anchor tag).
type Link struct {
	Href string
	Text string
}

// Parses an HTML document and returns a slice of Links.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	if n.Type == html.ElementNode && n != nil {
		fmt.Printf("%s<%s>\n", padding, n.Data)	
	}

	if n != nil && n.Type == html.TextNode {
		fmt.Println(padding, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding + "  ")
	}
}