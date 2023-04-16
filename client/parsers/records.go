package parsers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// parseRecordNode parses a record node.
func parseRecordNode(node *html.Node) (record models.Record, err error) {
	var c *html.Node

	for c = node.FirstChild; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			var parentId int
			parentId, err = strconv.Atoi(htmlquery.InnerText(c))

			if err != nil {
				err = &ErrParsing{
					`//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class] // parentId`,
					err,
				}
				return
			}

			record.ParentId = uint(parentId)
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			var recordId int
			recordId, err = strconv.Atoi(htmlquery.InnerText(c))

			if err != nil {
				err = &ErrParsing{
					`//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class] // recordId`,
					err,
				}
				return
			}

			rId := uint(recordId)
			record.Id = &rId
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if class := htmlquery.SelectAttr(c, "class"); len(class) == 0 {
			// pass
		} else if strings.Contains("dns_view_locked", class) {
			record.Domain = htmlquery.InnerText(c)

			if class == "dns_view_locked" {
				record.Locked = true
			}
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "center" {
			for d := c.FirstChild; d != nil; d = d.NextSibling {
				if d.Type != html.ElementNode || d.Data != "span" {
					// pass
				} else if data := htmlquery.SelectAttr(d, "data"); len(data) != 0 {
					record.RecordType = strings.ToUpper(data)
					break
				}
			}
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "left" {
			var recordTTL int
			recordTTL, err = strconv.Atoi(htmlquery.InnerText(c)) // To improve

			if err != nil {
				err = &ErrParsing{
					`//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class] // recordTTL`,
					err,
				}
				return
			}

			record.TTL = uint(recordTTL)
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "center" {
			p := htmlquery.InnerText(c)
			var priority int
			if priority, err = strconv.Atoi(p); err == nil {
				p := uint16(priority)
				record.Priority = &p
			} else if p != "-" {
				err = &ErrParsing{
					`//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class] // priority`,
					fmt.Errorf(`unknown priority value "%s"`, p),
				}
				return
			}
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") != "left" {
			// pass
		} else if data := htmlquery.SelectAttr(c, "data"); len(data) != 0 {
			record.Data = data
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			record.Dynamic, err = strconv.ParseBool(htmlquery.InnerText(c))

			if err != nil {
				err = &ErrParsing{
					`//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class] // dynamic`,
					err,
				}
				return
			}

			break
		}
	}

	return
}

// GetRecords returns the records from the HTML body.
func GetRecords(doc *html.Node) ([]models.Record, error) {
	q := `//div[@id="dns_main_content"]/table[@class="generictable"]`

	if table := htmlquery.FindOne(doc, q); table == nil {
		return nil, &ErrNotFound{q}
	}

	q = `//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class]`
	nodes := htmlquery.Find(doc, q)

	records := make([]models.Record, len(nodes))

	for i, node := range nodes {
		if record, err := parseRecordNode(node); err == nil {
			records[i] = record
		} else {
			return nil, err
		}
	}

	return records, nil
}
