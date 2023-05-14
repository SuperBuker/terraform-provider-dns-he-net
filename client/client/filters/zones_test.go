package filters_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/filters"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"

	"github.com/stretchr/testify/assert"
)

var domains = []models.Domain{
	{Id: 0, Domain: "a.example.com"},
	{Id: 1, Domain: "b.example.com"},
	{Id: 2, Domain: "c.example.com"},
	{Id: 3, Domain: "d.example.com"},
	{Id: 4, Domain: "e.example.com"},
}

func TestDomainById(t *testing.T) {
	for i := uint(0); i < 5; i++ {
		domain, ok := filters.DomainById(domains, i)
		assert.Equal(t, domains[i], domain)
		assert.True(t, ok)
	}

	domain, ok := filters.DomainById(domains, 5)
	assert.Equal(t, models.Domain{}, domain)
	assert.False(t, ok)
}

func TestDomainByName(t *testing.T) {
	for i := uint(0); i < 5; i++ {
		d := domains[i]
		domain, ok := filters.DomainByName(domains, d.Domain)
		assert.Equal(t, d, domain)
		assert.True(t, ok)
	}

	domain, ok := filters.DomainByName(domains, "")
	assert.Equal(t, models.Domain{}, domain)
	assert.False(t, ok)
}

func TestLatestDomain(t *testing.T) {
	domain, ok := filters.LatestDomain(domains)
	assert.Equal(t, domains[4], domain)
	assert.True(t, ok)

	domain, ok = filters.LatestDomain([]models.Domain{})
	assert.Equal(t, models.Domain{}, domain)
	assert.False(t, ok)
}
