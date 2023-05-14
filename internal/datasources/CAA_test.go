package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCAA(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_caa" "record-caa" {
					id = 5195537735
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "data", `0 issuewild ";"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "id", "5195537735"),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccCAAMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_caa" "record-caa" {
					id = 5195537735
					parent_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccCAAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_caa" "record-caa" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find CAA record"),
			},
		},
	})
}
