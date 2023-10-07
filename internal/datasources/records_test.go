package datasources_test

import (
	"fmt"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRecords(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: ProviderConfig +
					fmt.Sprintf(`data "dns-he-net_records" "test" {
					id = %d
				}`, Zone.ID),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					//resource.TestCheckResourceAttr("data.dns-he-net_records.test", "records.#", "104"), // TODO: enable

					// Verify placeholder attributes
					resource.TestCheckResourceAttr("data.dns-he-net_records.test", "id", toString(Zone.ID)),
				),
			},
		},
	})
}
