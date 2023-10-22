package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHINFO(t *testing.T) {
	record, ok := Records["HINFO"]
	if !ok {
		t.Skip("HINFO record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_hinfo" "record-hinfo" {
					id = 5195561437
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "domain", Zone.Sub("example-hinfo")),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "data", `"armv7 Linux"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_hinfo.record-hinfo", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccHINFOMissingZone(t *testing.T) {
	record, ok := Records["HINFO"]
	if !ok {
		t.Skip("HINFO record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_hinfo" "record-hinfo" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccHINFOMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_hinfo" "record-hinfo" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find HINFO record"),
			},
		},
	})
}
