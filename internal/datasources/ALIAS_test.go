package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccALIAS(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_alias" "record-alias" {
					id = 5195504740
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "domain", "example-alias.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "data", "dns-he-net.ovh"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "id", "5195504740"),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccALIASMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_alias" "record-alias" {
					id = 5195504740
					parent_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccALIASMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_alias" "record-alias" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find ALIAS record"),
			},
		},
	})
}
