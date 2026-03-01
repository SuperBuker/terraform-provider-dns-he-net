package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainZones(t *testing.T) {
	if Account.ID == "" {
		t.Skip("AccountID missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_domain_zones" "example" {
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_domain_zones.example", "zones.#", fmt.Sprint(domainZonesCount)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_domain_zones.example", "id", Account.ID),
				),
			},
		},
	})
}
