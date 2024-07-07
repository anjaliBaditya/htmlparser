package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
	Attrs map[string]string // Additional attributes
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

// buildLink constructs a Link from an HTML node.
func buildLink(n *html.Node) Link {
	var ret Link
	ret.Attrs = make(map[string]string)
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
		}
		ret.Attrs[attr.Key] = attr.Val
	}
	ret.Text = extractText(n)
	return ret
}

// extractText recursively extracts and concatenates text content from an HTML node.
func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += extractText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

// linkNodes recursively finds and returns all <a> nodes in the HTML document.
func linkNodes(n *html.Node) []*html.Node {
	var ret []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		ret = append(ret, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}

// FilterLinksByAttr filters links by a specified attribute and its value.
func FilterLinksByAttr(links []Link, attr, value string) []Link {
	var filtered []Link
	for _, link := range links {
		if v, ok := link.Attrs[attr]; ok && v == value {
			filtered = append(filtered, link)
		}
	}
	return filtered
}

// ExtractLinksText returns the text content of all links.
func ExtractLinksText(links []Link) []string {
	var texts []string
	for _, link := range links {
		texts = append(texts, link.Text)
	}
	return texts
}
