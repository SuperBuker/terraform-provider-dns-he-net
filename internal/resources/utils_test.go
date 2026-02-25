package resources_test

import (
	"fmt"

	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

var (
	// Capitalised variables are accessed by the entire the test package
	_resources     = test_cfg.Config.Resources
	ProviderConfig = _resources.Account.ProviderConfig("simple")
	Account        = _resources.Account
	ArpaZone       = _resources.ArpaZone
	DomainZone     = _resources.DomainZone
)

func getID(rawState map[string]string) string {
	return fmt.Sprintf("%s-%s", rawState["zone_id"], rawState["id"])
}

func importStateId(resourceName string) func(*terraform.State) (string, error) {
	return func(state *terraform.State) (string, error) {

		for _, m := range state.Modules {
			if len(m.Resources) > 0 {
				if v, ok := m.Resources[resourceName]; ok {
					rawState := v.Primary.Attributes
					return getID(rawState), nil
				}
			}
		}

		return "", fmt.Errorf("resource %q found", resourceName)
	}
}
