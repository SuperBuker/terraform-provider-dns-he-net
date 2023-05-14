package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSPF(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_spf" "record-spf" {
					id = 5195729389
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "data", `"v=spf1 include:_spf.example.com ~all"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "id", "5195729389"),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccSPFMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_spf" "record-spf" {
					id = 5195729389
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSPFMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_spf" "record-spf" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find SPF record"),
			},
		},
	})
}
