package resources_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccARecord(t *testing.T) {
	t.Parallel()

	domains := Zone.RandSub("example-%04d", 9999, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	password := randStringBytesMaskImprSrcSB(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Validate config
			// Must fail because the default dynamic value is false and data is not set
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 300
				}`, Zone.ID, domainInit),
				ExpectError: regexp.MustCompile("Invalid A record configuration"),
			},
			// Create and Read testing
			// Validates data default value by setting dynamic to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 300
					dynamic = true
				}`, Zone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "data", "0.0.0.0"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "dynamic", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_a.record-a",
				ImportStateIdFunc: importStateId("dns-he-net_a.record-a"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			// Updates ttl and domain
			// Sets dynamic to false and data to a known value
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "1.2.3.4"
				}`, Zone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "data", "1.2.3.4"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "dynamic", "false"),
				),
			},
			// Update and Read testing
			// Validates state continuity by setting dynamic to true and omitting data
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainUpdate, Zone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "data", "1.2.3.4"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Validates ddns update by retrieving the new data value
			{
				PreConfig: func() {
					// Force the ddns "external" update

					authObj, err := Account.Auth(auth.Simple)
					require.NoError(t, err)

					cli, err := client.NewClient(context.TODO(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
					require.NoError(t, err)

					assert.Equal(t, Account.ID, cli.GetAccount())

					ok, err := cli.DDNS().UpdateIP(context.TODO(), domainUpdate, password, "10.2.3.4")
					require.NoError(t, err)
					assert.True(t, ok)
				},
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainUpdate, Zone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "data", "10.2.3.4"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Validates forcing a data value with dynamic set to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_a" "record-a" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "0.0.0.0"
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, Zone.ID, domainUpdate, Zone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "zone_id", toString(Zone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "data", "0.0.0.0"),
					resource.TestCheckResourceAttr("dns-he-net_a.record-a", "dynamic", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
