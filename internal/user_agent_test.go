package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAgentString(t *testing.T) {
	expected := "HashiCorp Terraform/v0.14.0 (+https://www.terraform.io) terraform-provider-dns-he-net/v0.0.1 (+https://registry.terraform.io/providers/SuperBuker/dns-he-net)"
	assert.Equal(t, expected, UserAgentString(context.Background(), "v0.0.1", "v0.14.0"))

	add := "foo/bar"
	t.Setenv(uaEnvVar, add)

	expected = expected + " " + add
	assert.Equal(t, expected, UserAgentString(context.Background(), "v0.0.1", "v0.14.0"))
}
