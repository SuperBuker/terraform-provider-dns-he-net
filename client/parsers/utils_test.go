package parsers

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
)

var prefixToArpaZoneTests = []struct {
	subnet   string
	expected string
	valid    bool
}{
	{"2001:470:1f13:1::/64", "1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa", true},
	{"2001:470:1::/48", "1.0.0.0.0.7.4.0.1.0.0.2.ip6.arpa", true},
	{"::/0", "0.ip6.arpa", true},
	{"::", "0.ip6.arpa", false},
	{"invalid value", "", false},
	// TODO: These functions still accept invalid IPv6 entires like "2001:470:1::/129" or "127.0.0.1"
}

func TestPrefixToArpaZone(t *testing.T) {
	for i, test := range prefixToArpaZoneTests {
		prefix := models.NetworkPrefix{ID: uint(i), Value: test.subnet}
		arpaZone, err := NetworkPrefixToArpaZone(prefix)
		if test.valid {
			assert.Equal(t, uint(i), arpaZone.ID, "Unexpected Zone ID")
			assert.Equal(t, test.expected, arpaZone.Name, "Unexpected Zone Name")
		} else {
			assert.Error(t, err, "PrefixToArpaZone(%q) expected error", test.subnet)
		}
	}
}
