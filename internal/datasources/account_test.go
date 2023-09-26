package datasources_test

import (
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccount(t *testing.T) {
	accountID := os.Getenv("DNSHENET_ACCOUNT_ID")

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test_utils.ProviderConfig + `data "dns-he-net_account" "test" {}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify record attibutes
					resource.TestCheckResourceAttr("data.dns-he-net_account.test", "id", accountID),
				),
			},
		},
	})
}
