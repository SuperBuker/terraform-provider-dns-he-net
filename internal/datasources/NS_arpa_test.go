package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArpaNS(t *testing.T) {
	record, ok := ArpaZoneRecords["NS"]
	if !ok {
		t.Skip("NS record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_ns" "record-ns" {
					id = %d
					zone_id = %d
				}`, record.ID, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "data", record.Data),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "zone_id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}
