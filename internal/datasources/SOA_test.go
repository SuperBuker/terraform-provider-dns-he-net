package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSOA(t *testing.T) {
	record, ok := Records["SOA"]
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
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "ttl", "172800"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "mname", "ns1.he.net."),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "rname", "hostmaster.he.net."),
					//resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "serial", ""), # Updated constantly
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "refresh", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "retry", "7200"),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "expire", "3600000"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccSOAMissingZone(t *testing.T) {
	record, ok := Records["SOA"]
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
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSOAMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_soa" "record-soa" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find SOA record"),
			},
		},
	})
}
