package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSRV(t *testing.T) {
	record, ok := Records["SRV"]
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
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "domain", "_bofher._tcp.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "ttl", "28800"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "priority", "0"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "weight", "0"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "port", "22"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "target", Zone.Name),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "id", "5195753926"),
					resource.TestCheckResourceAttr("data.dns-he-net_srv.record-srv", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccSRVMissingZone(t *testing.T) {
	record, ok := Records["SRV"]
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
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find SRV record"),
			},
		},
	})
}
