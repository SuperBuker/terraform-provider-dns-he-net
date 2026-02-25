package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSOA(t *testing.T) {
	record, ok := DomainZoneRecords["SOA"]
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
				}`, record.ID, DomainZone.ID),
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
					resource.TestCheckResourceAttr("data.dns-he-net_soa.record-soa", "zone_id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccSOAMissingZone(t *testing.T) {
	record, ok := DomainZoneRecords["SOA"]
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
				}`, DomainZone.ID),
				ExpectError: regexp.MustCompile("Unable to find SOA record"),
			},
		},
	})
}
