package main

import (
	"flag"
	"fmt"

	. "linkparser.mmedic.com/m/v2/src/utils/file_reader"
	htmlparser "linkparser.mmedic.com/m/v2/src/utils/html_parser"
)

func main() {
	filename := flag.String("filename", "page.html", "Filename of the HTML file that needs to be parsed.")
	flag.Parse()

	fr := CreateFileReader()
	contents, err := fr.ReadFileAsString(*filename)

	if err != nil {
		panic(err)
	}

	hp := htmlparser.CreateHTMLParser()

	links, err := hp.GetLinks(contents)

	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}
