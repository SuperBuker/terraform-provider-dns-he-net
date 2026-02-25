package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetworkPrefix(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_network_prefix" "example" {
					id = %d
				}`, NetworkPrefix.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_network_prefix.example", "value", NetworkPrefix.Value),
					resource.TestCheckResourceAttr("data.dns-he-net_network_prefix.example", "enabled", fmt.Sprint(NetworkPrefix.Enabled)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_network_prefix.example", "id", fmt.Sprint(NetworkPrefix.ID)),
				),
			},
		},
	})
}

func TestAccMissingNetworkPrefix(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_network_prefix" "test" {
					id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to find network prefix"),
			},
		},
	})
}
