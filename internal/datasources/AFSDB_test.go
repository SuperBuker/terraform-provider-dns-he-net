package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAFSDB(t *testing.T) {
	record, ok := Records["AFSDB"]
	if !ok {
		t.Skip("AFSDB record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_afsdb" "record-afsdb" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "data", fmt.Sprintf("2 %s", Zone.Sub("green"))),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_afsdb.record-afsdb", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccAFSDBMissingZone(t *testing.T) {
	record, ok := Records["AFSDB"]
	if !ok {
		t.Skip("AFSDB record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_afsdb" "record-afsdb" {
				id = %d
				zone_id = 0
			}`, record.ID),
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
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_afsdb" "record-afsdb" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find AFSDB record"),
			},
		},
	})
}
