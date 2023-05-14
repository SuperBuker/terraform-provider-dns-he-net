package tfmodels

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModelCNAME(t *testing.T) {
	id := uint(1)

	exapected := models.CNAME{
		Id:     &id,
		ZoneID: 1,
		Domain: "example.com",
		TTL:    300,
		Data:   "new-example.com",
	}

	cname := CNAME{}

	require.NoError(t, cname.SetRecord(exapected))

	actual, err := cname.GetRecord()
	require.NoError(t, err)

	assert.Equal(t, exapected, actual)
}
