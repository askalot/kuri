package htmlutils

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func GetHTMLDocument(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	document, err := html.Parse(response.Body)
	if err != nil {
		return nil, err
	}

	return document, nil
}

func GetDOMNodeByTag(root *html.Node, tag string) (*html.Node, bool) {
	var traverse func(*html.Node) *html.Node

	traverse = func(node *html.Node) *html.Node {
		if node.Type == html.ElementNode && node.Data == tag {
			return node
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			if result := traverse(child); result != nil {
				return result
			}
		}

		return nil
	}

	return traverse(root), traverse(root) != nil
}

func GetHTMLTagTextFromURL(tag string, url string) (string, error) {
	tagNotFoundError := errors.New(fmt.Sprintf("Cannot find tag: %s", tag))

	root, err := GetHTMLDocument(url)
	if err != nil {
		return "", err
	}

	node, found := GetDOMNodeByTag(root, tag)
	if !found || node == nil || node.FirstChild == nil {
		return "", tagNotFoundError
	}

	return node.FirstChild.Data, nil
}
