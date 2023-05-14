package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSOA(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_soa" "record-soa" {
					id = 5182379278
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "ttl", "172800"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "mname", "ns1.he.net."),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "rname", "hostmaster.he.net."),
					//resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "serial", ""), # Updated constantly
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "refresh", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "retry", "7200"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "expire", "3600000"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "id", "5182379278"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccSOAMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_soa" "record-soa" {
					id = 5182379278
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSOAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_soa" "record-soa" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find SOA record"),
			},
		},
	})
}
