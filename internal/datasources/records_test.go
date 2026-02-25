package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDomainZoneRecords(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_records" "test" {
					id = %d
				}`, DomainZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_records.test", "records.#", fmt.Sprint(DomainZone.RecordCount)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_records.test", "id", fmt.Sprint(DomainZone.ID)),
				),
			},
		},
	})
}

func TestAccArpaZoneRecords(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_records" "test" {
					id = %d
				}`, ArpaZone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attributes
					resource.TestCheckResourceAttr("data.dns-he-net_records.test", "records.#", fmt.Sprint(ArpaZone.RecordCount)),

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_records.test", "id", fmt.Sprint(ArpaZone.ID)),
				),
			},
		},
	})
}
