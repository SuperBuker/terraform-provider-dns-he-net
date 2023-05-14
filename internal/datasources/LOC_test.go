package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccLOC(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_loc" "record-loc" {
					id = 5195584706
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_loc.record-loc", "domain", "example-loc.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_loc.record-loc", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_loc.record-loc", "data", "40 27 53.86104 N 3 39 2.59092 W 712.8m 0.00m 0.00m 0.00m"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_loc.record-loc", "id", "5195584706"),
					resource.TestCheckResourceAttr("data.dns-he-net_loc.record-loc", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccLOCMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_loc" "record-loc" {
					id = 5195584706
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccLOCMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_loc" "record-loc" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find LOC record"),
			},
		},
	})
}
