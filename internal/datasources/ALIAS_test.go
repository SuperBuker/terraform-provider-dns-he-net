package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccALIAS(t *testing.T) {
	record, ok := DomainZoneRecords["ALIAS"]
	if !ok {
		t.Skip("ALIAS record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_alias" "record-alias" {
					id = %d
					zone_id = %d
				}`, record.ID, DomainZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "data", record.Data),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "zone_id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccALIASMissingZone(t *testing.T) {
	record, ok := DomainZoneRecords["ALIAS"]
	if !ok {
		t.Skip("ALIAS record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_alias" "record-alias" {
					id = %d
					zone_id = 0
				}`, record.ID),
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
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_alias" "record-alias" {
					id = 0
					zone_id = %d
				}`, DomainZone.ID),
				ExpectError: regexp.MustCompile("Unable to find ALIAS record"),
			},
		},
	})
}
