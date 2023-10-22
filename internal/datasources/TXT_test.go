package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccTXT(t *testing.T) {
	record, ok := Records["TXT"]
	if !ok {
		t.Skip("TXT record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig + `data "dns-he-net_txt" "record-txt" {
					id = 5195711991
					zone_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "domain", Zone.Sub("bofher")),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "ttl", "300"),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "data", `"Just for the record"`),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "dynamic", "true"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_txt.record-txt", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccTXTMissingZone(t *testing.T) {
	record, ok := Records["TXT"]
	if !ok {
		t.Skip("TXT record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_txt" "record-txt" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccTXTMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_txt" "record-txt" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find TXT record"),
			},
		},
	})
}
