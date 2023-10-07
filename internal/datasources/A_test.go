package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccARecord(t *testing.T) {
	record, ok := Records["A"]
	if !ok {
		t.Skip("A record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_a" "record-a" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "domain", Zone.Sub("example-a")),
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "data", "0.0.0.0"),
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "dynamic", "true"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccAMissingZone(t *testing.T) {
	record, ok := Records["A"]
	if !ok {
		t.Skip("A record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_a" "record-a" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_a" "record-a" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find A record"),
			},
		},
	})
}
