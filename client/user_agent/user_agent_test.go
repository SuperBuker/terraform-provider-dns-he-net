package user_agent_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/user_agent"
	"github.com/stretchr/testify/assert"
)

func TestUserAgentProduct(t *testing.T) {
	p := user_agent.UserAgentProduct{
		Name:    "HashiCorp Terraform",
		Version: "v0.14.0",
		Comment: "+https://www.terraform.io",
	}

	expected := "HashiCorp Terraform/v0.14.0 (+https://www.terraform.io)"

	assert.Equal(t, expected, p.String())
}

func TestUserAgentProducts(t *testing.T) {
	p := user_agent.UserAgentProducts{
		user_agent.UserAgentProduct{
			Name:    "HashiCorp Terraform",
			Version: "v0.14.0",
			Comment: "+https://www.terraform.io",
		},
		user_agent.UserAgentProduct{
			Name:    "terraform-provider-dns-he-net",
			Version: "0.0.1",
			Comment: "+https://registry.terraform.io/providers/SuperBuker/dns-he-net",
		},
	}

	expected := "HashiCorp Terraform/v0.14.0 (+https://www.terraform.io) terraform-provider-dns-he-net/0.0.1 (+https://registry.terraform.io/providers/SuperBuker/dns-he-net)"

	assert.Equal(t, expected, p.String())
}
