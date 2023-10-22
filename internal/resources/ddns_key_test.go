package resources_test

import (
	"context"
	"fmt"
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

	domains := Zone.RandSubs("example-%04d", 10000, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	password := randStringBytesMaskImprSrcSB(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			// Validate defaults
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainInit, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Update and Read testing
			{
				PreConfig: func() {
					// Force the ddns key to be updated externally
					authObj, err := Account.Auth(auth.Simple)
					require.NoError(t, err)

					cli, err := client.NewClient(context.Background(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
					require.NoError(t, err)

					assert.Equal(t, Account.ID, cli.GetAccount())

					// Makes auth fail when validating the expected key, triggering an update
					anotherPassword := randStringBytesMaskImprSrcSB(16)

					ddnsKey := models.DDNSKey{
						Domain: domainInit,
						ZoneID: Zone.ID,
						Key:    anotherPassword,
					}

					_, err = cli.SetDDNSKey(context.Background(), ddnsKey)
					require.NoError(t, err)
				},
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainInit, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Update and Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "id", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_ddnskey.ddnskey", "key", password),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
