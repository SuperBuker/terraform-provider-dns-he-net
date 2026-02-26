package resources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/resources"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccAAAARecord(t *testing.T) {
	t.Parallel()

	domains := DomainZone.RandSubs("example-%04d", 10000, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	password := resources.GenerateRandomString(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Validate config
			// Must fail because the default dynamic value is false and data is not set
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 300
				}`, DomainZone.ID, domainInit),
				ExpectError: regexp.MustCompile("Invalid AAAA record configuration"),
			},
			// Create and Read testing
			// Validates data default value by setting dynamic to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 300
					dynamic = true
				}`, DomainZone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "data", "::"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "dynamic", "true"),
				),
			},
			/*{
				Check TestCheckFunc
			},*/
			// ImportState testing
			{
				ResourceName:      "dns-he-net_aaaa.record-aaaa",
				ImportStateIdFunc: importStateId("dns-he-net_aaaa.record-aaaa"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			// Updates ttl and domain
			// Sets dynamic to false and data to a known value
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "::1"
				}`, DomainZone.ID, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "data", "::1"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "dynamic", "false"),
				),
			},
			// Update and Read testing
			// Validates state continuity by setting dynamic to true and omitting data
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, DomainZone.ID, domainUpdate, DomainZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "data", "::1"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Validates ddns update by retrieving the new data value
			{
				PreConfig: func() {
					// Force the ddns "external" update
					authObj, err := Account.Auth(auth.Simple)
					require.NoError(t, err)

					cli, err := client.NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false))
					require.NoError(t, err)

					assert.Equal(t, Account.ID, cli.GetAccount())

					ok, err := cli.DDNS().UpdateIP(t.Context(), domainUpdate, password, "::2")
					require.NoError(t, err)
					assert.True(t, ok)
				},
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, DomainZone.ID, domainUpdate, DomainZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "data", "::2"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Validates forcing a data value with dynamic set to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_aaaa" "record-aaaa" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "::"
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, DomainZone.ID, domainUpdate, DomainZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "zone_id", fmt.Sprint(DomainZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "data", "::"),
					resource.TestCheckResourceAttr("dns-he-net_aaaa.record-aaaa", "dynamic", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
		//CheckDestroy: testCheckDestroy, //TODO
	})
}
