package result

import (
	"bytes"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html"
)

// ResultX repurposes go-resty Request.Result to persist a copy
// of the parsed HTML body. This way we prevent the parsing to
// be performed on each step of the HTML analysis.
type ResultX struct {
	HTML   *html.Node
	Result interface{}
}

// Init initialises the go-resty Request.Result with a ResultX contaning the
// parsed HTML body and the expected result, if any. Returns an error if the
// parsing fails
func Init(resp *resty.Response) (err error) {
	body := resp.Body()
	res := resp.Result()

	var doc *html.Node
	doc, err = htmlquery.Parse(bytes.NewReader(body))

	if err != nil {
		return err
	}

	if utils.IsNil(res) {
		res = ResultX{HTML: doc}
	} else if resX, ok := resp.Request.Result.(ResultX); ok {
		// Case of request retrial
		res = ResultX{HTML: doc, Result: resX.Result}
	} else {
		res = ResultX{HTML: doc, Result: res}
	}

	resp.Request.Result = res
	return
}

// Returns the parsed HTML body from the go-resty Response.
// If the Request.Result contains a ResultX, it returns the HTML body from it,
// otherwise it parses the body and returns the result.
func Body(resp *resty.Response) *html.Node {
	res, ok := resp.Result().(ResultX)

	if !ok {
		// Not nice but ensures backwards compatibility
		doc, _ := htmlquery.Parse(bytes.NewReader(resp.Body()))
		return doc
	}

	return res.HTML
}

// Returns the Result() the go-resty Response.
// If the Request.Result contains a ResultX, it returns the inner Result.
func Result(resp *resty.Response) interface{} {
	res, ok := resp.Result().(ResultX)

	if !ok {
		return resp.Result()
	}

	return res.Result
}
