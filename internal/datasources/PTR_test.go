package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPTR(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ptr" "record-ptr" {
					id = 5195612976
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "domain", "example-ptr.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "data", "dns-he-net.ovh"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "id", "5195612976"),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccPTRMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ptr" "record-ptr" {
					id = 5195612976
					zone_id = 0
				}`,
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccPTRMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_ptr" "record-ptr" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find PTR record"),
			},
		},
	})
}
