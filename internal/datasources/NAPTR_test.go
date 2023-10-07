package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNAPTR(t *testing.T) {
	record, ok := Records["NAPTR"]
	if !ok {
		t.Skip("NAPTR record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_naptr" "record-naptr" {
					id = 5195590349
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "data", fmt.Sprintf(`100 10 "S" "SIP+D2U" "!^.*$!sip:bofher@%s!" _sip._udp.%s.`, Zone.Name, Zone.Name)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccNAPTRMissingZone(t *testing.T) {
	record, ok := Records["NAPTR"]
	if !ok {
		t.Skip("NAPTR record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_naptr" "record-naptr" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccNAPTRMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_naptr" "record-naptr" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find NAPTR record"),
			},
		},
	})
}
