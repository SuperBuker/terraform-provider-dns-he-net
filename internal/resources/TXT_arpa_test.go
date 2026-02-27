package resources_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccArpaTXTRecord(t *testing.T) {
	t.Parallel()

	domains := ArpaZone.RandArpaSubs(16, 2)
	domainInit := domains[0]
	domainUpdate := domains[1]

	data := `"` + utils.GenerateRandomString(600) + `"`
	data2 := `"` + utils.GenerateRandomString(100) + `"` // The API doesn't support large TXT records

	password := utils.GenerateRandomString(16)

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Validate config
			// Must fail because the default dynamic value is false and data is not set
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 300
				}`, ArpaZone.ID, domainInit),
				ExpectError: regexp.MustCompile("Invalid TXT record configuration"),
			},
			// Create and Read testing
			// Validates data default value by setting dynamic to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 300
					dynamic = true
				}`, ArpaZone.ID, domainInit),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "domain", domainInit),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "ttl", "300"),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "data", `""`),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "dynamic", "true"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "dns-he-net_txt.record-txt",
				ImportStateIdFunc: importStateId("dns-he-net_txt.record-txt"),
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			// Updates ttl and domain
			// Sets dynamic to false and data to a known value
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = %q
				}`, ArpaZone.ID, domainUpdate, data),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "data", data),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "dynamic", "false"),
				),
			},
			// Update and Read testing
			// Validates state continuity by setting dynamic to true and omitting data
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, ArpaZone.ID, domainUpdate, ArpaZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "data", data),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "dynamic", "true"),
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

					ok, err := cli.DDNS().UpdateTXT(t.Context(), domainUpdate, password, data2[1:len(data2)-1])
					require.NoError(t, err)
					assert.True(t, ok)
				},
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 600
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, ArpaZone.ID, domainUpdate, ArpaZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "data", data2),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "dynamic", "true"),
				),
			},
			// Update and Read testing
			// Validates forcing a data value with dynamic set to true
			{
				Config: ProviderConfig +
					fmt.Sprintf(`resource "dns-he-net_txt" "record-txt" {
					zone_id = %d
					domain = %q
					ttl = 600
					data = "\"some data\""
					dynamic = true
				}

				resource "dns-he-net_ddnskey" "ddnskey" {
					zone_id = %d
					domain = %q
					key = %q
				}`, ArpaZone.ID, domainUpdate, ArpaZone.ID, domainUpdate, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "zone_id", fmt.Sprint(ArpaZone.ID)),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "domain", domainUpdate),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "ttl", "600"),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "data", `"some data"`),
					resource.TestCheckResourceAttr("dns-he-net_txt.record-txt", "dynamic", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
