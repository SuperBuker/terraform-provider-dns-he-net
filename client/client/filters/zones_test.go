package filters_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/stretchr/testify/assert"
)

var zones = []models.Zone{
	{ID: 0, Name: "a.example.com"},
	{ID: 1, Name: "b.example.com"},
	{ID: 2, Name: "c.example.com"},
	{ID: 3, Name: "d.example.com"},
	{ID: 4, Name: "e.example.com"},
}

func TestZoneById(t *testing.T) {
	for i := uint(0); i < 5; i++ {
		zone, ok := filters.ZoneById(zones, i)
		assert.Equal(t, zones[i], zone)
		assert.True(t, ok)
	}

	zone, ok := filters.ZoneById(zones, 5)
	assert.Equal(t, models.Zone{}, zone)
	assert.False(t, ok)
}

func TestZoneByName(t *testing.T) {
	for i := uint(0); i < 5; i++ {
		d := zones[i]
		zone, ok := filters.ZoneByName(zones, d.Name)
		assert.Equal(t, d, zone)
		assert.True(t, ok)
	}

	zone, ok := filters.ZoneByName(zones, "")
	assert.Equal(t, models.Zone{}, zone)
	assert.False(t, ok)
}

func TestLatestZone(t *testing.T) {
	zone, ok := filters.LatestZone(zones)
	assert.Equal(t, zones[4], zone)
	assert.True(t, ok)

	zone, ok = filters.LatestZone([]models.Zone{})
	assert.Equal(t, models.Zone{}, zone)
	assert.False(t, ok)
}
