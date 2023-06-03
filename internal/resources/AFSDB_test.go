package resources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAFSDBRecord(t *testing.T) {
	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	password := randStringBytesMaskImprSrcSB(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Validate config
			// Must fail because the default dynamic value is false and data is not set
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 300
				}`, domainInit),
				ExpectError: regexp.MustCompile("Invalid AFSDB record configuration"),
			},
			// Create and Read testing
			// Validates data default value by setting dynamic to true
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					dynamic = true
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "1 afsdb.example.com"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "dynamic", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_afsdb.record-afsdb",
				ImportStateIdFunc: importStateId("dns-he-net_afsdb.record-afsdb"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			// Updates ttl and domain
			// Sets dynamic to false and data to a known value
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "1 green.dns-he-net.eu.org"
				}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "1 green.dns-he-net.eu.org"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "dynamic", "false"),
				),
			},
			// Update and Read testing
			// Validates state continuity by setting dynamic to true and ommiting data
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					domain = %q
					zone_id = 1091256
					key = %q
				}`, domainUpdate, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "1 green.dns-he-net.eu.org"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Update and Read testing
			// Validates forcing a data value with dynamic set to true
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "1 green.dns-he-net.eu.org"
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					domain = %q
					zone_id = 1091256
					key = %q
				}`, domainUpdate, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "1 green.dns-he-net.eu.org"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "dynamic", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
