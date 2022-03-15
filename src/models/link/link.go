package link

import "strings"

type Link struct {
	href string
	text string
}

func CreateEmptyLink() *Link {
	return &Link{}
}

func (l *Link) SetHref(href string) {
	l.href = strings.TrimSpace(href)
}

func (l *Link) SetText(text string) {
	l.text = strings.TrimSpace(text)
}

func (l *Link) GetHref() string {
	return l.href
}

func (l *Link) GetText() string {
	return l.text
}
