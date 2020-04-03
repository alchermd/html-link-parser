package main

import (
	"strings"
	"github.com/alchermd/linkparser"
)

func main() {
	exampleHtml := `<a href="/foo">bar</a>`
	r := strings.NewReader(exampleHtml)
	_, err := linkparser.Parse(r)
	if err != nil {
		panic(err)
	}
}