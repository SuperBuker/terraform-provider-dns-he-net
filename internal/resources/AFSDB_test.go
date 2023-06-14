package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAFSDBRecord(t *testing.T) {
	t.Parallel()
	
	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			// Validates data default value by setting dynamic to true
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_afsdb" "record-afsdb" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "1 green.dns-he-net.eu.org"
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "1 green.dns-he-net.eu.org"),
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
					data = "2 blue.dns-he-net.eu.org"
				}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_afsdb.record-afsdb", "data", "2 blue.dns-he-net.eu.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
