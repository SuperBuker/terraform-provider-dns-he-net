package filters_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/stretchr/testify/assert"
)

var networkPrefixes = []models.NetworkPrefix{
	{ID: 0, Value: "2001:470:1f13:343::/64", Enabled: true},
	{ID: 1, Value: "2001:470:1f13:344::/64", Enabled: true},
	{ID: 2, Value: "2001:470:1f13:345::/64", Enabled: false},
}

func TestNetworkPrefixById(t *testing.T) {
	for _, netPref := range networkPrefixes {
		p, ok := filters.NetworkPrefixById(networkPrefixes, netPref.ID)
		assert.True(t, ok)
		assert.Equal(t, netPref, p)
	}

	// missing ID
	p, ok := filters.NetworkPrefixById(networkPrefixes, 3)
	assert.False(t, ok)
	assert.Equal(t, models.NetworkPrefix{}, p)

	// empty slice
	p, ok = filters.NetworkPrefixById([]models.NetworkPrefix{}, 0)
	assert.False(t, ok)
	assert.Equal(t, models.NetworkPrefix{}, p)
}

func TestNetworkPrefixByValue(t *testing.T) {
	for _, netPref := range networkPrefixes {
		p, ok := filters.NetworkPrefixByValue(networkPrefixes, netPref.Value)
		assert.True(t, ok)
		assert.Equal(t, netPref, p)
	}

	// missing ID
	p, ok := filters.NetworkPrefixByValue(networkPrefixes, "")
	assert.False(t, ok)
	assert.Equal(t, models.NetworkPrefix{}, p)

	// empty slice
	p, ok = filters.NetworkPrefixByValue([]models.NetworkPrefix{}, "2001:470:1f13:343::/64")
	assert.False(t, ok)
	assert.Equal(t, models.NetworkPrefix{}, p)
}
