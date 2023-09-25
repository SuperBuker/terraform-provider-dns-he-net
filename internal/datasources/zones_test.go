package datasources_test

import (
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccZones(t *testing.T) {
	t.Parallel()
	accountID := os.Getenv("DNSHENET_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_zones" "example" {
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_zones.example", "zones.#", "3"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_zones.example", "id", accountID),
				),
			},
		},
	})
}
