package params_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client/params"

	"github.com/stretchr/testify/assert"
)

func TestRecordCreate(t *testing.T) {
	x := map[string]string{
		"menu":                  "edit_zone",
		"hosted_dns_editzone":   "1",
		"hosted_dns_editrecord": "Submit",
	}

	m := map[string]string{}
	n := params.RecordCreate(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}

func TestRecordUpdate(t *testing.T) {
	x := map[string]string{
		"menu":                  "edit_zone",
		"hosted_dns_editzone":   "1",
		"hosted_dns_editrecord": "Update",
	}

	m := map[string]string{}
	n := params.RecordUpdate(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}

func TestRecordDelete(t *testing.T) {
	x := map[string]string{
		"menu":                  "edit_zone",
		"hosted_dns_editzone":   "1",
		"hosted_dns_delconfirm": "delete",
		"hosted_dns_delrecord":  "1",
	}

	m := map[string]string{}
	n := params.RecordDelete(m)

	assert.Equal(t, x, m)
	assert.Equal(t, x, n)
	assert.Equal(t, m, n)

	m["test"] = "test"
	assert.Equal(t, m, n)
	// It's effectively the same map
}
