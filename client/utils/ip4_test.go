package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/stretchr/testify/assert"
)

// WIP...
var ipv4SubnetToPTRTests = []struct {
	subnet        string
	expected_arpa string
}{
	{"192.0.2.1/24", "1.2.0.192.in-addr.arpa"},
	{"198.51.100.0/24", "0.100.51.198.in-addr.arpa"},
	{"203.0.113.0/24", "0.113.0.203.in-addr.arpa"},
}

func TestIPv4AddrToPTR(t *testing.T) {
	for _, test := range ipv4SubnetToPTRTests {
		ipNet, err := utils.ParseIPNet(test.subnet)
		if err != nil {
			t.Errorf("ParseIPNet(%q) returned error: %v", test.subnet, err)
			continue
		}

		arpaZone := utils.IPv4AddrToPTR(ipNet.IP)
		assert.Equal(t, test.expected_arpa, arpaZone, "IPv4AddrToPTR(%q) = %q; want %q", test.subnet, arpaZone, test.expected_arpa)
	}
}
