package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSPFRecord(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_spf" "record-spf" {
					zone_id = 1091256
					domain = %q
					ttl = 300
					data = "\"v=spf1 -all\""
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "data", `"v=spf1 -all"`),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_spf.record-spf",
				ImportStateIdFunc: importStateId("dns-he-net_spf.record-spf"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_spf" "record-spf" {
					zone_id = 1091256
					domain = %q
					ttl = 600
					data = "\"v=spf1 a:example.com ~all\""
			}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_spf.record-spf", "data", `"v=spf1 a:example.com ~all"`),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
