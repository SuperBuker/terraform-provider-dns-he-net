package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHINFORecord(t *testing.T) {
	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_hinfo" "record-hinfo" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "\"armv7 Linux\""
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "data", `"armv7 Linux"`),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_hinfo.record-hinfo",
				ImportStateIdFunc: importStateId("dns-he-net_hinfo.record-hinfo"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_hinfo" "record-hinfo" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "\"amd64 Linux\""
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_hinfo.record-hinfo", "data", `"amd64 Linux"`),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}