package params_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"
	"github.com/stretchr/testify/assert"
)

func TestDDNSKeySet(t *testing.T) {
	x := map[string]string{
		"menu":                "edit_zone",
		"hosted_dns_editzone": "1",
		"generate_key":        "Submit",
	}

	m := map[string]string{}
	n := params.DDNSKeySet(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}
