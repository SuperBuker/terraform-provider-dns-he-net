package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCAA(t *testing.T) {
	record, ok := Records["CAA"]
	if !ok {
		t.Skip("CAA record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_caa" "record-caa" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "data", `0 issuewild ";"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccCAAMissingZone(t *testing.T) {
	record, ok := Records["CAA"]
	if !ok {
		t.Skip("CAA record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_caa" "record-caa" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccCAAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_caa" "record-caa" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find CAA record"),
			},
		},
	})
}
