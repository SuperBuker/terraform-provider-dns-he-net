package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSSHFPRecord(t *testing.T) {
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
					fmt.Sprintf(`resource "dns-he-net_sshfp" "record-sshfp" {
					zone_id = %d
					domain = %q
					ttl = 300
					data = "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"
				}`, Zone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "data", "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_sshfp.record-sshfp",
				ImportStateIdFunc: importStateId("dns-he-net_sshfp.record-sshfp"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_sshfp" "record-sshfp" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "4 2 123456789Abcdef67890123456789abcdef67890123456789abcdef123456789"
			}`, Zone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_sshfp.record-sshfp", "data", "4 2 123456789Abcdef67890123456789abcdef67890123456789abcdef123456789"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
