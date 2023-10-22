package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCNAME(t *testing.T) {
	record, ok := Records["CNAME"]
	if !ok {
		t.Skip("CNAME record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_cname" "record-cname" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "domain", Zone.Sub("example-cname")),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "data", "example.com"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_cname.record-cname", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccCNAMEMissingZone(t *testing.T) {
	record, ok := Records["CNAME"]
	if !ok {
		t.Skip("CNAME record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_cname" "record-cname" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccCNAMEMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_cname" "record-cname" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find CNAME record"),
			},
		},
	})
}
