package test_cfg

import (
	"fmt"
	"math/rand"
)

type uniqueRand struct {
	bound     uint
	generated map[int]struct{}
}

func (u *uniqueRand) Int() int {
	for {
		var i int
		if u.bound > 0 {
			i = rand.Intn(int(u.bound))
		} else {
			i = rand.Int()
		}

		if _, ok := u.generated[i]; !ok {
			u.generated[i] = struct{}{}
			return i
		}
	}
}

func newUniqueRand(bound uint) *uniqueRand {
	return &uniqueRand{bound: bound, generated: make(map[int]struct{})}
}

func generateSubDomains(template string, bound int, count int) []string {
	if bound < count && bound > 0 { // bound == 0 means no bound
		panic("bound must be greater than len")
	}

	generator := newUniqueRand(uint(bound))

	domains := make([]string, count)
	for j := 0; j < count; j++ {
		domains[j] = fmt.Sprintf(template, generator.Int())
	}
	return domains
}
