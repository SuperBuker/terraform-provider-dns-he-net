package models_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
)

func TestDDNSKey(t *testing.T) {
	ddns := models.DDNSKey{
		Domain: "example.com",
		ZoneID: 123,
		Key:    "secret",
	}

	t.Run("serialisation", func(t *testing.T) {
		serialised := ddns.Serialise()

		expected := map[string]string{
			"hosted_dns_zoneid": "123",
			"Name":              "example.com",
			"Key":               "secret",
			"Key2":              "secret",
		}
		assert.Equal(t, expected, serialised)

	})

	t.Run("refs", func(t *testing.T) {
		refs := ddns.Refs()

		expected := map[string]string{
			"hosted_dns_zoneid": "123",
			"Name":              "example.com",
		}
		assert.Equal(t, expected, refs)
	})

	t.Run("getters", func(t *testing.T) {
		assert.Equal(t, "example.com", ddns.GetDomain())
		assert.Equal(t, uint(123), ddns.GetZoneID())
	})
}
