package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLOCRecord(t *testing.T) {
	t.Parallel()
	
	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_loc" "record-loc" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "data", "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_loc.record-loc",
				ImportStateIdFunc: importStateId("dns-he-net_loc.record-loc"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_loc" "record-loc" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_loc.record-loc", "data", "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
