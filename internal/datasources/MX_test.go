package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMX(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_mx" "record-mx" {
					id = 5195493159
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "ttl", "3600"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "priority", "1"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "data", "mx.example.com"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "id", "5195493159"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccMXMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_mx" "record-mx" {
					id = 5195493159
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccMXMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_mx" "record-mx" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find MX record"),
			},
		},
	})
}
