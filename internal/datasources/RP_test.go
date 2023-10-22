package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRP(t *testing.T) {
	record, ok := Records["RP"]
	if !ok {
		t.Skip("RP record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_rp" "record-rp" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "data", fmt.Sprintf("bofher.%s bofher.%s", Zone.Name, Zone.Name)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_rp.record-rp", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccRPMissingZone(t *testing.T) {
	record, ok := Records["RP"]
	if !ok {
		t.Skip("RP record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_rp" "record-rp" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccRPMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_rp" "record-rp" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find RP record"),
			},
		},
	})
}
