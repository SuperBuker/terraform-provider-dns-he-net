package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArpaPTR(t *testing.T) {
	record, ok := ArpaZoneRecords["PTR"]
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
				}`, record.ID, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "data", record.Data),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_ptr.record-ptr", "zone_id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}
