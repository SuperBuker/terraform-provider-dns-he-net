package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPTRRecord(t *testing.T) {
	t.Parallel()

	domains := Zone.RandSub("example-%04d", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_ptr" "record-ptr" {
					zone_id = %d
					domain = %q
					ttl = 300
					data = "example.com"
				}`, Zone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "data", "example.com"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_ptr.record-ptr",
				ImportStateIdFunc: importStateId("dns-he-net_ptr.record-ptr"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_ptr" "record-ptr" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "example.io"
			}`, Zone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_ptr.record-ptr", "data", "example.io"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
