package utils_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/utils"
	"github.com/stretchr/testify/assert"
)

var ipv6NetTests = []struct {
	ipNet string
	valid bool
}{
	{"192.168.0.0/32", true},       // IPv4
	{"2001:470:1f13:1::/64", true}, // IPv6
	{"0.0.0.0/0", true},            // Global IPv4
	{"::/0", true},                 // Global IPv6
	{"256.168.0.0/32", false},      // Invalid IPv4
	{"::/129", false},              // Invalid IPv6 mask length
	{"::/abc", false},              // Invalid IPv6 mask format
	{"::", false},                  // Missing mask
	{"invalid value", false},       // Invalid format
}

func TestParseIPNet(t *testing.T) {
	for _, test := range ipv6NetTests {
		ipNet, err := utils.ParseIPNet(test.ipNet)
		if test.valid {
			assert.NotNil(t, ipNet, "ParseIPNet(%q) = nil; want non-nil", test.ipNet)
			assert.NoError(t, err, "ParseIPNet(%q) unexpected error: %v", test.ipNet, err)
		} else {
			assert.Error(t, err, "ParseIPNet(%q) expected error", test.ipNet)
		}
	}
}
