package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNS(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ns" "record-ns" {
					id = 5182379279
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "ttl", "172800"),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "data", "ns1.he.net"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "id", "5182379279"),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccNSMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ns" "record-ns" {
					id = 5182379279
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccNSMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ns" "record-ns" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find NS record"),
			},
		},
	})
}
