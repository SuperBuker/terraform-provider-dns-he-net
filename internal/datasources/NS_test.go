package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNS(t *testing.T) {
	record, ok := Records["NS"]
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
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "ttl", "172800"),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "data", "ns1.he.net"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_ns.record-ns", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccNSMissingZone(t *testing.T) {
	record, ok := Records["NS"]
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
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccNSMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_ns" "record-ns" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find NS record"),
			},
		},
	})
}
