package parsers

import (
	"bytes"

	"github.com/antchfx/htmlquery"
)

func GetAccount(data []byte) (string, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		return "", err
	}

	q := `//form[@name="remove_domain"]/input[@name="account"]`
	node := htmlquery.FindOne(doc, q)
	// TODO: Check if node is nil

	return htmlquery.SelectAttr(node, "value"), nil
}
