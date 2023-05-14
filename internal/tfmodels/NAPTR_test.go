package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelNAPTR(t *testing.T) {
	id := uint(1)

	expected := models.NAPTR{
		Id:       &id,
		ParentId: 1,
		Domain:   "example.com",
		TTL:      86400,
		Data:     `100 10 "S" "SIP+D2U" "!^.*$!sip:example.com!" _sip._udp.example.com.`,
	}

	naptr := NAPTR{}

	require.NoError(t, naptr.SetRecord(expected))

	actual, err := naptr.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
