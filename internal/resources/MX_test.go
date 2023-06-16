package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMXRecord(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_mx" "record-mx" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					priority = 10
					data = "mx.example.com"
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "data", "mx.example.com"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_mx.record-mx",
				ImportStateIdFunc: importStateId("dns-he-net_mx.record-mx"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_mx" "record-mx" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					priority = 20
					data = "mx.example.io"
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "data", "mx.example.io"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
