package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArpaSOA(t *testing.T) {
	record, ok := ArpaZoneRecords["SOA"]
	if !ok {
		t.Skip("SOA record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_soa" "record-soa" {
					id = %d
					zone_id = %d
				}`, record.ID, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "mname", record.ExtraArgs["mname"]),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "rname", record.ExtraArgs["rname"]),
					//resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "serial", ""), # Updated constantly
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "refresh", record.ExtraArgs["refresh"]),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "retry", record.ExtraArgs["retry"]),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "expire", record.ExtraArgs["expire"]),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "zone_id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}
