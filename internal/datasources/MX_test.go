package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccMX(t *testing.T) {
	record, ok := Records["MX"]
	if !ok {
		t.Skip("MX record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_mx" "record-mx" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "ttl", "3600"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "priority", "1"),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "data", "mx.example.com"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_mx.record-mx", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccMXMissingZone(t *testing.T) {
	record, ok := Records["MX"]
	if !ok {
		t.Skip("MX record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_mx" "record-mx" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccMXMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_mx" "record-mx" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find MX record"),
			},
		},
	})
}
