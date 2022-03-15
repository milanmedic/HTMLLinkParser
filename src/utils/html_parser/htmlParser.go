package htmlparser

import (
	"bytes"
	"strings"

	. "linkparser.mmedic.com/m/v2/src/models/link"

	"golang.org/x/net/html"
)

type HTMLParser struct{}

func CreateHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

func (hp *HTMLParser) GetLinks(contents string) (map[string]Link, error) {

	htmlReader := strings.NewReader(contents)
	tokenizer := html.NewTokenizer(htmlReader)

	var links map[string]Link = make(map[string]Link)
	linkFound := false
	var link *Link

	for {
		tokenizerToken := tokenizer.Next()

		switch tokenizerToken {
		case html.ErrorToken:
			return nil, tokenizer.Err()
		case html.TextToken:
			handleTextToken(links, link, &linkFound, tokenizer)
		case html.StartTagToken:
			link = handleStartTagToken(links, &linkFound, tokenizer)
		case html.EndTagToken:
			tagName, _ := tokenizer.TagName()
			if string(tagName) == "html" {
				return links, nil
			}
		}
	}
}

func handleTextToken(links map[string]Link, link *Link, linkFound *bool, tokenizer *html.Tokenizer) {
	if *linkFound {
		if !bytes.Equal(tokenizer.Text(), []byte(html.CommentToken.String())) {
			link.SetText(string(tokenizer.Text()))
			links[link.GetHref()] = *link
			*linkFound = false
		}
	}
}

func handleStartTagToken(links map[string]Link, linkFound *bool, tokenizer *html.Tokenizer) *Link {
	tagName, _ := tokenizer.TagName()
	if string(tagName) == "a" {
		link := CreateEmptyLink()
		*linkFound = true
		_, attrValue, _ := tokenizer.TagAttr()
		link.SetHref(string(attrValue))
		return link
	}
	return nil
}
