package resources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSRVRecord(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_srv" "record-srv" {
					zone_id = 1091256
					domain = "_bofher._tcp.dns-he-net.eu.org"
					ttl = 300
					port = 80
					target = %q
				}`, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "domain", "_bofher._tcp.dns-he-net.eu.org"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "priority", "0"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "weight", "0"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "port", "80"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "target", domainInit),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_srv.record-srv",
				ImportStateIdFunc: importStateId("dns-he-net_srv.record-srv"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_srv" "record-srv" {
					zone_id = 1091256
					domain = "_bofher._udp.dns-he-net.eu.org"
					ttl = 600
					priority = 10
					weight = 10
					port = 80
					target = %q
				}`, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "domain", "_bofher._udp.dns-he-net.eu.org"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "priority", "10"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "weight", "10"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "port", "80"),
					resource.TestCheckResourceAttr("dns-he-net_srv.record-srv", "target", domainUpdate),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
