package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNAPTR(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_naptr" "record-naptr" {
					id = 5195590349
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "domain", "dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "data", `100 10 "S" "SIP+D2U" "!^.*$!sip:bofher@dns-he-net.ovh!" _sip._udp.dns-he-net.ovh.`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "id", "5195590349"),
					resource.TestCheckResourceAttr("data.dns-he-net_naptr.record-naptr", "zone_id", "1093397"),
				),
			},
		},
	})
}

func TestAccNAPTRMissingZone(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_naptr" "record-naptr" {
					id = 5195590349
					zone_id = 0
				}`,
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
				Config: test_utils.ProviderConfig + `data "dns-he-net_naptr" "record-naptr" {
					id = 0
					zone_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find NAPTR record"),
			},
		},
	})
}
