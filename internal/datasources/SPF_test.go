package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSPF(t *testing.T) {
	record, ok := Records["SPF"]
	if !ok {
		t.Skip("SPF record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_spf" "record-spf" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "domain", Zone.Name),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "data", `"v=spf1 include:_spf.example.com ~all"`),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_spf.record-spf", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccSPFMissingZone(t *testing.T) {
	record, ok := Records["SPF"]
	if !ok {
		t.Skip("SPF record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_spf" "record-spf" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSPFMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_spf" "record-spf" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find SPF record"),
			},
		},
	})
}
