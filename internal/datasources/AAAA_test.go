package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAAAARecord(t *testing.T) {
	record, ok := DomainZoneRecords["AAAA"]
	if !ok {
		t.Skip("AAAA record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_aaaa" "record-aaaa" {
					id = %d
					zone_id = %d
				}`, record.ID, DomainZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "data", record.Data),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "dynamic", fmt.Sprint(record.Dynamic)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccAAAAMissingZone(t *testing.T) {
	record, ok := DomainZoneRecords["AAAA"]
	if !ok {
		t.Skip("AAAA record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_aaaa" "record-aaaa" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccAAAAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_aaaa" "record-aaaa" {
					id = 0
					zone_id = %d
				}`, DomainZone.ID),
				ExpectError: regexp.MustCompile("Unable to find AAAA record"),
			},
		},
	})
}
