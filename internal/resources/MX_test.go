package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMXRecord(t *testing.T) {
	t.Parallel()

	domains := Zone.RandSubs("example-%04d", 10000, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_mx" "record-mx" {
					zone_id = %d
					domain = %q
					ttl = 300
					priority = 10
					data = "mx.example.com"
				}`, Zone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "zone_id", toString(Zone.ID)),
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
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_mx" "record-mx" {
					zone_id = %d
					domain = %q
					ttl = 600
					priority = 20
					data = "mx.example.io"
			}`, Zone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_mx.record-mx", "data", "mx.example.io"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
