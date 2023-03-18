package parsers

import (
	"bytes"
	"log"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"

	"github.com/antchfx/htmlquery"
)

func LoginStatus(data []byte) (auth.Status, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		log.Fatal(err)
		return auth.Unknown, err
	}

	q := `//form[@name="login"]`
	node := htmlquery.FindOne(doc, q)

	if node == nil {
		return auth.Ok, nil
	}

	q = `//input[@id="tfacode"]`
	node = htmlquery.FindOne(doc, q)

	if node == nil {
		return auth.NoAuth, nil
	}

	return auth.OTP, nil
}
