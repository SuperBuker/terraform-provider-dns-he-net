package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArpaZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_arpa_zone" "example" {
					zone_id = %d
				}`, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_arpa_zone.example", "name", ArpaZone.Name),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_arpa_zone.example", "zone_id", fmt.Sprint(ArpaZone.ID)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_arpa_zone.example", "zone_id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}

func TestAccMissingArpaZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_arpa_zone" "test" {
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to find arpa zone"),
			},
		},
	})
}
