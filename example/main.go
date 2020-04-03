package main

import (
	"strings"
	"github.com/alchermd/linkparser"
	"fmt"
)

func main() {
	exampleHtml := `
<body>
<div>
    <h1>Example Data</h1>
    <p>Lorem ipsum <a href="/lorem">dolor</a></p>
    <p><a href="/info">More information about our <span>services</span> <!-- Some comment --></a></p>
</div>
</body>
</html>
`
	r := strings.NewReader(exampleHtml)
	links, err := linkparser.Parse(r)
	if err != nil {
		panic(err)
	}
	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}