package resources_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDDNSKey(t *testing.T) {
	t.Parallel()

	domains := generateSubDomains("example-%04d.dns-he-net.eu.org", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	password := randStringBytesMaskImprSrcSB(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			// Validate defaults
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					domain = %q
					zone_id = 1091256
					key = %q
				}`, domainInit, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// Force the ddns key to be updated externally
					user := os.Getenv("DNSHENET_USER")
					password := os.Getenv("DNSHENET_PASSWD")
					otp := os.Getenv("DNSHENET_OTP")

					authObj, err := auth.NewAuth(user, password, otp, auth.Simple)
					require.NoError(t, err)

					cli, err := client.NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
					require.NoError(t, err)

					assert.Equal(t, "v6643873d8c41428.97783691", cli.GetAccount())

					// Makes auth fail when validating the expected key, triggering an update
					anotherPassword := randStringBytesMaskImprSrcSB(16)

					ddnsKey := models.DDNSKey{
						Domain: domainInit,
						ZoneID: 1091256,
						Key:    anotherPassword,
					}

					_, err = cli.SetDDNSKey(context.TODO(), ddnsKey)
					require.NoError(t, err)
				},
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					domain = %q
					zone_id = 1091256
					key = %q
				}`, domainInit, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Update and Read testing
			{
				Config: test_utils.ProviderConfig + fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					domain = %q
					zone_id = 1091256
					key = %q
				}`, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", "1091256"),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
