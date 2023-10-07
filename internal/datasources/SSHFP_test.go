package datasources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSSHFP(t *testing.T) {
	record, ok := Records["SSHFP"]
	if !ok {
		t.Skip("SSHFP record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_sshfp" "record-sshfp" {
					id = %d
					zone_id = %d
				}`, record.ID, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "domain", Zone.Sub("example-sshfp")),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "data", "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "id", toString(record.ID)),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "zone_id", toString(Zone.ID)),
				),
			},
		},
	})
}

func TestAccSSHFPMissingZone(t *testing.T) {
	record, ok := Records["SSHFP"]
	if !ok {
		t.Skip("SSHFP record missing in config")
	}

	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_sshfp" "record-sshfp" {
					id = %d
					zone_id = 0
				}`, record.ID),
				ExpectError: regexp.MustCompile("Unable to fetch DNS records"),
			},
		},
	})
}

func TestAccSSHFPMissingRecord(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_sshfp" "record-sshfp" {
					id = 0
					zone_id = %d
				}`, Zone.ID),
				ExpectError: regexp.MustCompile("Unable to find SSHFP record"),
			},
		},
	})
}
