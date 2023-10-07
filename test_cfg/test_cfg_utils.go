package test_cfg

import (
	"fmt"
	"math/rand"
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
