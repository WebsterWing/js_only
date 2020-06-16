// Given HTML input on stdin, print all of the inline
// javascript in the document
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Document couln't be parsed as html")
	}
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			if n.FirstChild == nil {
				return
			}
			if n.FirstChild.Type == html.TextNode {
				fmt.Println("//-------------------------------------")
				fmt.Println(n.FirstChild.Data)
			}
		}
	})
}

// Recursively call pre on each node in the document
func forEachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre)
	}
}
