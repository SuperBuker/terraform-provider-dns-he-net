package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNAPTRRecord(t *testing.T) {
	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_naptr" "record-naptr" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "example.com"
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "data", "example.com"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_naptr.record-naptr",
				ImportStateIdFunc: importStateId("dns-he-net_naptr.record-naptr"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_naptr" "record-naptr" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "example.io"
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "data", "example.io"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
