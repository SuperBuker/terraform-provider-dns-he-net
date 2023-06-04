package resources_test

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

type uniqueRand struct {
	size      uint
	generated map[int]struct{}
}

func (u *uniqueRand) Int() int {
	for {
		var i int
		if u.size > 0 {
			i = rand.Intn(int(u.size))
		} else {
			i = rand.Int()
		}

		if _, ok := u.generated[i]; !ok {
			u.generated[i] = struct{}{}
			return i
		}
	}
}

func newUniqueRand(size uint) *uniqueRand {
	return &uniqueRand{size: size, generated: make(map[int]struct{})}
}

func generateSubDomains(template string, size int, len int) []string {
	// TODO: Display alert if size is lower than len.

	generator := newUniqueRand(uint(size))

	domains := make([]string, len)
	for j := 0; j < len; j++ {
		domains[j] = fmt.Sprintf(template, generator.Int())
	}
	return domains
}

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
