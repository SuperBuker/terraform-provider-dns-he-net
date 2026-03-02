package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCAA(t *testing.T) {
	record, ok := DomainZoneRecords["CAA"]
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
				}`, record.ID, DomainZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "data", record.Data),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_caa.record-caa", "zone_id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccCAAMissingZone(t *testing.T) {
	record, ok := DomainZoneRecords["CAA"]
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
				}`, DomainZone.ID),
				ExpectError: regexp.MustCompile("Unable to find CAA record"),
			},
		},
	})
}
