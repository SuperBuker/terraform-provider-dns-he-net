package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelSOA(t *testing.T) {
	id := uint(1)

	expected := models.SOA{
		Id:       &id,
		ParentId: 1,
		Domain:   "example.com",
		TTL:      172800,
		MName:    "ns1.he.net.",
		RName:    "hostmaster.he.net.",
		Serial:   2023050324,
		Refresh:  86400,
		Retry:    7200,
		Expire:   3600000,
	}

	soa := SOA{}

	require.NoError(t, soa.SetRecord(expected))

	actual, err := soa.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
