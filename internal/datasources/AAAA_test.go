package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAAAARecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_aaaa" "record-aaaa" {
					id = 5195455723
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "domain", "example-aaaa.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "data", "::"),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "dynamic", "true"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "id", "5195455723"),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccAAAAMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_aaaa" "record-aaaa" {
					id = 5195455723
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccAAAAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_aaaa" "record-aaaa" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find AAAA record"),
			},
		},
	})
}
