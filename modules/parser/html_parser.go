package parser

import (
	"bytes"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type HTMLParser struct{}

func NewHTMLParser() *HTMLParser {
	return &HTMLParser{}
}

func (parser *HTMLParser) ReadFile(path string) ([]string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return parser.ParseBytes(file)
}

func (parser *HTMLParser) ParseBytes(raw []byte) ([]string, error) {
	doc, err := html.Parse(bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}

	var lines []string
	var extractParagraphs func(*html.Node)
	extractParagraphs = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "p" {
			var buf bytes.Buffer
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				html.Render(&buf, c)
			}
			text := strings.TrimSpace(buf.String())
			if text != "" {
				lines = append(lines, text)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractParagraphs(c)
		}
	}

	extractParagraphs(doc)

	return lines, nil
}

func (parser *HTMLParser) WriteFile(path string, data []string) error {
	raw, err := parser.SerializeBytes(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

func (parser *HTMLParser) SerializeBytes(data []string) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("<html><body>\n")
	for _, line := range data {
		buf.WriteString("<p>" + line + "</p>\n")
	}
	buf.WriteString("</body></html>")
	return buf.Bytes(), nil
}
