package datasources_test

import (
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSSHFP(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_sshfp" "record-sshfp" {
					id = 5195770791
					parent_id = 1093397
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "domain", "example-sshfp.dns-he-net.ovh"),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "ttl", "86400"),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "data", "4 2 123456789abcdef67890123456789abcdef67890123456789abcdef123456789"),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "id", "5195770791"),
					resource.TestCheckResourceAttr("data.dns-he-net_sshfp.record-sshfp", "parent_id", "1093397"),
				),
			},
		},
	})
}

func TestAccSSHFPMissingDomain(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_sshfp" "record-sshfp" {
					id = 5195770791
					parent_id = 0
				}`,
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
				Config: test_utils.ProviderConfig + `data "dns-he-net_sshfp" "record-sshfp" {
					id = 0
					parent_id = 1093397
				}`,
				ExpectError: regexp.MustCompile("Unable to find SSHFP record"),
			},
		},
	})
}
