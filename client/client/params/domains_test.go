package params_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"

	"github.com/stretchr/testify/assert"
)

func TestDomainCreate(t *testing.T) {
	x := map[string]string{
		"action":   "add_zone",
		"retmain:": "0",
		"submit":   "Add Domain!",
	}

	m := map[string]string{}
	n := params.DomainCreate(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}

func TestDomainDelete(t *testing.T) {
	x := map[string]string{
		"remove_domain": "1",
	}

	m := map[string]string{}
	n := params.DomainDelete(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}
