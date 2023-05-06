package result_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/result"
	"github.com/antchfx/htmlquery"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func setup(t *testing.T) *resty.Client {
	client := resty.New()

	// block all HTTP requests
	httpmock.ActivateNonDefault(client.GetClient())

	return client
}

func teardown(t *testing.T) {
	httpmock.DeactivateAndReset()
}

func TestResultX(t *testing.T) {
	client := setup(t)
	defer teardown(t)

	data := `<html><head><title>Example Domain</title></head><body>some data</body></html>`
	responder := httpmock.NewStringResponder(200, data)
	url := "https://example.com"
	httpmock.RegisterResponder("GET", url, responder)

	// Parse html
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = result.Init(resp)
		}
		return
	})

	t.Run("Without Result", func(t *testing.T) {
		resp, err := client.R().Get(url)
		require.NoError(t, err)

		body := resp.Body()

		// Compare parsing results
		var doc *html.Node
		doc, err = htmlquery.Parse(bytes.NewReader(body))
		require.NoError(t, err)

		assert.Equal(t, doc, result.Body(resp))

		// Validate parsing
		q := `//body`
		node := htmlquery.FindOne(doc, q)
		assert.Equal(t, "some data", htmlquery.InnerText(node))

		res, ok := resp.Result().(result.ResultX)
		require.True(t, ok)

		// Compare result output
		assert.Nil(t, res.Result)
		assert.Equal(t, res.Result, result.Result(resp))
	})

	t.Run("With Result", func(t *testing.T) {
		resp, err := client.R().SetResult(struct{}{}).Get(url)
		require.NoError(t, err)

		body := resp.Body()

		// Compare parsing results
		var doc *html.Node
		doc, err = htmlquery.Parse(bytes.NewReader(body))
		require.NoError(t, err)

		assert.Equal(t, doc, result.Body(resp))

		// Validate parsing
		q := `//body`
		node := htmlquery.FindOne(doc, q)
		assert.Equal(t, "some data", htmlquery.InnerText(node))

		res, ok := resp.Result().(result.ResultX)
		require.True(t, ok)

		// Compare result output
		assert.NotNil(t, res.Result)
		assert.Equal(t, res.Result, result.Result(resp))
	})
}

func TestNoResultX(t *testing.T) {
	client := setup(t)
	defer teardown(t)

	data := `<html><head><title>Example Domain</title></head><body>some data</body></html>`
	responder := httpmock.NewStringResponder(200, data)
	url := "https://example.com"
	httpmock.RegisterResponder("GET", url, responder)

	resp, err := client.R().Get(url)
	require.NoError(t, err)

	body := resp.Body()

	// Compare parsing results
	var doc *html.Node
	doc, err = htmlquery.Parse(bytes.NewReader(body))
	require.NoError(t, err)

	assert.Equal(t, doc, result.Body(resp))

	// Validate parsing
	q := `//body`
	node := htmlquery.FindOne(doc, q)
	assert.Equal(t, "some data", htmlquery.InnerText(node))

	_, ok := resp.Result().(result.ResultX)
	require.False(t, ok)

	// Compare result output
	assert.Nil(t, resp.Result())
	assert.Equal(t, resp.Result(), result.Result(resp))
}

func TestResultXRetry(t *testing.T) {
	// This test is required because it should be safe to initialise twice ResultX
	// on the same request. This is because it will be reinitialised on a retry.

	client := setup(t)
	defer teardown(t)

	data := `<html><head><title>Example Domain</title></head><body>some data</body></html>`
	responder := httpmock.NewStringResponder(200, data)
	url := "https://example.com"
	httpmock.RegisterResponder("GET", url, responder)

	// Parse html
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) (err error) {
		if resp.StatusCode() == 200 {
			err = errors.Join(result.Init(resp), result.Init(resp))
		}
		return
	})

	t.Run("Without Result", func(t *testing.T) {
		resp, err := client.R().Get(url)
		require.NoError(t, err)

		body := resp.Body()

		// Compare parsing results
		var doc *html.Node
		doc, err = htmlquery.Parse(bytes.NewReader(body))
		require.NoError(t, err)

		assert.Equal(t, doc, result.Body(resp))

		// Validate parsing
		q := `//body`
		node := htmlquery.FindOne(doc, q)
		assert.Equal(t, "some data", htmlquery.InnerText(node))

		res, ok := resp.Result().(result.ResultX)
		require.True(t, ok)

		// Compare result output
		assert.Nil(t, res.Result)
		assert.Equal(t, res.Result, result.Result(resp))
	})

	t.Run("With Result", func(t *testing.T) {
		resp, err := client.R().SetResult(struct{}{}).Get(url)
		require.NoError(t, err)

		body := resp.Body()

		// Compare parsing results
		var doc *html.Node
		doc, err = htmlquery.Parse(bytes.NewReader(body))
		require.NoError(t, err)

		assert.Equal(t, doc, result.Body(resp))

		// Validate parsing
		q := `//body`
		node := htmlquery.FindOne(doc, q)
		assert.Equal(t, "some data", htmlquery.InnerText(node))

		res, ok := resp.Result().(result.ResultX)
		require.True(t, ok)

		// Compare result output
		assert.NotNil(t, res.Result)
		assert.Equal(t, res.Result, result.Result(resp))
	})
}
