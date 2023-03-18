package parsers

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parseRecordNode(node *html.Node) (record models.Record) { // missing error
	var c *html.Node

	for c = node.FirstChild; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			parentId, _ := strconv.Atoi(htmlquery.InnerText(c)) // To improve
			record.ParentId = uint(parentId)
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			recordId, _ := strconv.Atoi(htmlquery.InnerText(c)) // To improve
			rId := uint(recordId)
			record.Id = &rId
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
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
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "center" {
			for d := c.FirstChild; d != nil; d = d.NextSibling {
				if d.Type != html.ElementNode {
					// pass
				} else if d.Data != "span" {
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
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "left" {
			//fmt.Println("ttl")
			recordTTL, _ := strconv.Atoi(htmlquery.InnerText(c)) // To improve
			record.TTL = uint(recordTTL)
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") == "center" {
			//fmt.Println("priority", htmlquery.InnerText(c))
			p := htmlquery.InnerText(c)
			if priority, err := strconv.Atoi(p); err == nil {
				p := uint16(priority)
				record.Priority = &p // To improve
			} else if p != "-" {
				// this is an error
				fmt.Println(err)
			}
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "align") != "left" {
			// pass
		} else if data := htmlquery.SelectAttr(c, "data"); len(data) != 0 {
			//fmt.Println("data", data)
			record.Data = data
			break
		}
	}

	for c = c.NextSibling; ; c = c.NextSibling {
		if c == nil {
			return
		} else if c.Type != html.ElementNode {
			// pass
		} else if c.Data != "td" {
			// pass
		} else if htmlquery.SelectAttr(c, "class") == "hidden" {
			//fmt.Println("dynamic", htmlquery.InnerText(c))
			record.Dynamic, _ = strconv.ParseBool(htmlquery.InnerText(c))
			break
		}
	}

	return
}

func GetRecords(data []byte) ([]models.Record, error) {
	doc, err := htmlquery.Parse(bytes.NewReader(data))

	if err != nil {
		return nil, err
	}

	q := `//div[@id="dns_main_content"]/table[@class="generictable"]/tbody/tr[@class]`
	nodes := htmlquery.Find(doc, q)

	records := make([]models.Record, len(nodes))

	for i, node := range nodes {
		record := parseRecordNode(node)
		records[i] = record
	}

	return records, nil
}
