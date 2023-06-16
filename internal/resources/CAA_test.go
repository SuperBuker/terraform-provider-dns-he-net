package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCAARecord(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_caa" "record-caa" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "0 iodef \"bofher@dns-he-net.eu.org\""
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "data", `0 iodef "bofher@dns-he-net.eu.org"`),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_caa.record-caa",
				ImportStateIdFunc: importStateId("dns-he-net_caa.record-caa"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_caa" "record-caa" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "0 issuewild \";\""
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_caa.record-caa", "data", `0 issuewild ";"`),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
