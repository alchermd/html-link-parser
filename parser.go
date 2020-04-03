package linkparser

import (
	"golang.org/x/net/html"
	"io"
	"strings"
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
	var links []Link
	nodes := getLinkNodes(doc)
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// Builds a Link from an HTML Node
func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

// Parses an HTML Node and finds all the anchor tags from its children
func getLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getLinkNodes(c)...)
	}
	return ret
}

// Build a string from an HTML nodes, recursively adding its children's text as well
func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}
