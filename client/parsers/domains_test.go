package parsers_test

import (
	"io/ioutil"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var domains = []models.Domain{
	{Id: 1234567, Domain: "example.com"},
}

func TestDomains(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data, err := ioutil.ReadFile("../testing_data/main.html")
		require.NoError(t, err)

		_domains, err := parsers.GetDomains(data)
		require.NoError(t, err)

		for i, domain := range _domains {
			assert.Equal(t, domains[i], domain)
		}
	})
}
