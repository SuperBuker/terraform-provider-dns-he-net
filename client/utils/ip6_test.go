package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/stretchr/testify/assert"
)

var ipv6SubnetToPTRTests = []struct {
	subnet        string
	expected_arpa string
	expected_ptr  string
}{
	{"2001:470:1f13:1::/64", "1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa", "0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa"},
	{"2001:470:1::/48", "1.0.0.0.0.7.4.0.1.0.0.2.ip6.arpa", "0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.1.0.0.0.0.7.4.0.1.0.0.2.ip6.arpa"},
	{"::/0", "0.ip6.arpa", "0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.ip6.arpa"},
}

func TestIPv6AddrToPTR(t *testing.T) {
	for _, test := range ipv6SubnetToPTRTests {
		ipNet, err := utils.ParseIPNet(test.subnet)
		if err != nil {
			t.Errorf("ParseIPNet(%q) returned error: %v", test.subnet, err)
			continue
		}

		arpaZone := utils.IPv6AddrToPTR(ipNet.IP, false)
		assert.Equal(t, test.expected_arpa, arpaZone, "IPv6AddrToPTR(%q) = %q; want %q", test.subnet, arpaZone, test.expected_arpa)

		ptr := utils.IPv6AddrToPTR(ipNet.IP, true)
		assert.Equal(t, test.expected_ptr, ptr, "IPv6AddrToPTR(%q) = %q; want %q", test.subnet, ptr, test.expected_ptr)
	}
}

func TestIPv6SubnetToArpaZone(t *testing.T) {
	t.Run("Valid subnets", func(t *testing.T) {
		for _, test := range ipv6SubnetToPTRTests {
			arpaZone, err := utils.IPv6SubnetToArpaZone(test.subnet)
			assert.Equal(t, test.expected_arpa, arpaZone, "IPv6SubnetToArpaZone(%q) = %q; want %q", test.subnet, arpaZone, test.expected_arpa)
			assert.NoError(t, err)
		}
	})

	t.Run("Invalid subnets", func(t *testing.T) {
		invalidSubnets := []string{
			"2001:470:1f13:1::/129",
			"2001:470:1f13:1::/64/64",
			"192.168.0.1/24",    // IPv4 subnet
			"2001:470:1f13:1::", // Missing mask
			"invalid",
		}
		for _, subnet := range invalidSubnets {
			_, err := utils.IPv6SubnetToArpaZone(subnet)
			assert.Error(t, err, "IPv6SubnetToArpaZone(%q) should return an error", subnet)
		}
	})
}
