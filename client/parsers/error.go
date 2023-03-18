package parsers

import (
	"bytes"

	"github.com/antchfx/htmlquery"
)

func ParseError(data []byte) (string, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		return "", err
	}

	q := `//div[@id="dns_err"]`
	node := htmlquery.FindOne(doc, q)

	if node != nil {
		return htmlquery.InnerText(node), nil
	}

	return "", nil
}
