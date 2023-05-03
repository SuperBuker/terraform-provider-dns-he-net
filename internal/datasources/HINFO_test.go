package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHINFO(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_hinfo" "record-hinfo" {
					id = 5195561437
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "domain", "example-hinfo.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "data", `"armv7 Linux"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "id", "5195561437"),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccHINFOMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_hinfo" "record-hinfo" {
					id = 5195561437
					parent_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccHINFOMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_hinfo" "record-hinfo" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find HINFO record"),
			},
		},
	})
}
