package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSRV(t *testing.T) {
	record, ok := DomainZoneRecords["SRV"]
	if !ok {
		t.Skip("SRV record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_srv" "record-srv" {
					id = %d
					zone_id = %d
				}`, record.ID, DomainZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "domain", record.Domain),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "ttl", fmt.Sprint(record.TTL)),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "priority", record.ExtraArgs["priority"]),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "weight", record.ExtraArgs["weight"]),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "port", record.ExtraArgs["port"]),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "target", DomainZone.Name),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "id", fmt.Sprint(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "zone_id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccSRVMissingZone(t *testing.T) {
	record, ok := DomainZoneRecords["SRV"]
	if !ok {
		t.Skip("SRV record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_srv" "record-srv" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSRVMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_srv" "record-srv" {
					id = 0
					zone_id = %d
				}`, DomainZone.ID),
				ExpectError: regexp.MustCompile("Unable to find SRV record"),
			},
		},
	})
}
