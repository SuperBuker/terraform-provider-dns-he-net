package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCNAME(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_cname" "record-cname" {
					id = 5195464830
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "domain", "example-cname.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "data", "example.com"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "id", "5195464830"),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccCNAMEMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_cname" "record-cname" {
					id = 5195464830
					parent_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccCNAMEMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_cname" "record-cname" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find CNAME record"),
			},
		},
	})
}
