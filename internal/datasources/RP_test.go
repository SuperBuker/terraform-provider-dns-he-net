package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRP(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_rp" "record-rp" {
					id = 5195714295
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "data", "bofher.dns-he-net.ovh bofher.dns-he-net.ovh"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "id", "5195714295"),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccRPMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_rp" "record-rp" {
					id = 5195714295
					parent_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccRPMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_rp" "record-rp" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find RP record"),
			},
		},
	})
}
