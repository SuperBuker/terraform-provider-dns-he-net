package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRPRecord(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_rp" "record-rp" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "bofher.dns-he-net.eu.org bofher.dns-he-net.eu.org"
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "data", "bofher.dns-he-net.eu.org bofher.dns-he-net.eu.org"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_rp.record-rp",
				ImportStateIdFunc: importStateId("dns-he-net_rp.record-rp"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_rp" "record-rp" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "bofher2.dns-he-net.eu.org bofher2.dns-he-net.eu.org"
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_rp.record-rp", "data", "bofher2.dns-he-net.eu.org bofher2.dns-he-net.eu.org"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
