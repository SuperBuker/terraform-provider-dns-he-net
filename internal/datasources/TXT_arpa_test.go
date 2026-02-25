package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccArpaTXT(t *testing.T) {
	record, ok := ArpaZoneRecords["TXT"]
	if !ok {
		t.Skip("TXT record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_txt" "record-txt" {
					id = %d
					zone_id = %d
				}`, record.ID, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "data", record.Data),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "dynamic", "false"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}
