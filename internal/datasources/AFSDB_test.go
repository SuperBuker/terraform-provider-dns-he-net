package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAFSDB(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_afsdb" "record-afsdb" {
					id = 5195520341
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "data", "2 green.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "dynamic", "true"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "id", "5195520341"),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccAFSDBMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_afsdb" "record-afsdb" {
					id = 5195520341
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccAFSDBMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_afsdb" "record-afsdb" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find AFSDB record"),
			},
		},
	})
}
