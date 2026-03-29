package parser

import (
	"golang.org/x/net/html"
	"strings"
)

func Parse(htmlStr string) (*html.Node, error) {
	return html.Parse(strings.NewReader(htmlStr))
}

func ExtractText(n *html.Node, out *[]string) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			*out = append(*out, text)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ExtractText(c, out)
	}
}
