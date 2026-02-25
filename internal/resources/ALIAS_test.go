package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccALIASRecord(t *testing.T) {
	t.Parallel()

	domains := DomainZone.RandSubs("example-%04d", 10000, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_alias" "record-alias" {
					zone_id = %d
					domain = %q
					ttl = 300
					data = %q
				}`, DomainZone.ID, domainInit, DomainZone.Name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "data", DomainZone.Name),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_alias.record-alias",
				ImportStateIdFunc: importStateId("dns-he-net_alias.record-alias"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_alias" "record-alias" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = %q
				}`, DomainZone.ID, domainUpdate, DomainZone.Name),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_alias.record-alias", "data", DomainZone.Name),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
