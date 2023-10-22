package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPTR(t *testing.T) {
	record, ok := Records["PTR"]
	if !ok {
		t.Skip("PTR record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_ptr" "record-ptr" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "domain", Zone.Sub("example-ptr")),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "data", Zone.Name),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccPTRMissingZone(t *testing.T) {
	record, ok := Records["PTR"]
	if !ok {
		t.Skip("PTR record missing in config")
	}

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_ptr" "record-ptr" {
					id = %d
					zone_id = 0
				}`, record.ID),
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
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_ptr" "record-ptr" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find PTR record"),
			},
		},
	})
}
