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
			var zoneID int
			zoneID, err = strconv.Atoi(htmlquery.InnerText(c))

			if err != nil {
				err = errParsingNode(recordQ, "zoneID", err)
				return
			}

			record.ZoneID = uint(zoneID)
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode || c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			var recordID int
			recordID, err = strconv.Atoi(htmlquery.InnerText(c))

			if err != nil {
				err = errParsingNode(recordQ, "recordID", err)
				return
			}

			rID := uint(recordID)
			record.ID = &rID
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
				err = errParsingNode(recordQ, "recordTTL", err)
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
				err = errParsingNode(recordQ, "priority", fmt.Errorf("unknown priority value %q", p))
				return
			} else {
				// Clean the error
				err = nil
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
			record.Data = htmlquery.SelectAttr(c, "data")
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
				err = errParsingNode(recordQ, "dynamic", err)
				return
			}

			break
		}
	}

	return
}

// GetRecords returns the records from the HTML body.
func GetRecords(doc *html.Node) ([]models.Record, error) {
	if table := htmlquery.FindOne(doc, recordsTableQ); table == nil {
		return nil, &ErrNotFound{recordsTableQ}
	}

	nodes := htmlquery.Find(doc, recordQ)
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
