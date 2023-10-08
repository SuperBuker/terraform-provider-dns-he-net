package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNAPTRRecord(t *testing.T) {
	t.Parallel()

	domains := Zone.RandSubs("example-%04d", 10000, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_naptr" "record-naptr" {
					zone_id = %d
					domain = %q
					ttl = 300
					data = "100 10 \"S\" \"SIP+D2U\" \"!^.*$!sip:bofher@dns-he-net.eu.org!\" _sip._udp.dns-he-net.eu.org."
				}`, Zone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "data", "100 10 \"S\" \"SIP+D2U\" \"!^.*$!sip:bofher@dns-he-net.eu.org!\" _sip._udp.dns-he-net.eu.org."),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_naptr.record-naptr",
				ImportStateIdFunc: importStateId("dns-he-net_naptr.record-naptr"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_naptr" "record-naptr" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "100 10 \"S\" \"SIP+D2T\" \"!^.*$!sip:bofher@dns-he-net.eu.org!\" _sip._tcp.dns-he-net.eu.org."
			}`, Zone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_naptr.record-naptr", "data", "100 10 \"S\" \"SIP+D2T\" \"!^.*$!sip:bofher@dns-he-net.eu.org!\" _sip._tcp.dns-he-net.eu.org."),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
