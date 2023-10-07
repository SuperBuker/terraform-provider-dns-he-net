package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccALIAS(t *testing.T) {
	record, ok := Records["ALIAS"]
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
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "domain", Zone.Sub("example-alias")),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "data", Zone.Name),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_alias.record-alias", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccALIASMissingZone(t *testing.T) {
	record, ok := Records["ALIAS"]
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
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find ALIAS record"),
			},
		},
	})
}
