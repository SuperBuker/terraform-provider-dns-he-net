package test_cfg_test

import (
	"os"
	"strings"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestZoneCfg(t *testing.T) {
	c := test_cfg.ZoneCfg{
		Name: "test.com",
	}

	assert.Equal(t, "sub.test.com", c.Sub("sub"))

	// Upper bound is exclusive, only 0 is generated
	assert.Equal(t, []string{"000.test.com"}, c.RandSubs("%03d", 1, 1))

	assert.Len(t, c.RandSubs("%d", 0, 1000), 1000)

	assert.PanicsWithValue(t, "bound must be greater than len", func() {
		c.RandSubs("%d", 1, 2) // bound < len
	})
}

func TestConfig(t *testing.T) {
	// Load the test configuration file
	config_path := os.Getenv("DNSHENET_TEST_CONFIG_PATH")
	if config_path == "" {
		t.Skip("DNSHENET_TEST_CONFIG_PATH is not set")
	}

	c := test_cfg.TestCfg{}
	require.NoError(t, c.Load(config_path), "This shouldn't fail")

	assert.Equal(t, test_cfg.Config, c)

	require.Error(t, c.Load(string(os.PathSeparator)), "This should fail")
}

func TestAccountCfg(t *testing.T) {
	c := test_cfg.AccountCfg{
		User:     "user",
		Password: "password",
		OTP:      "otp",
		ID:       "id",
	}

	assert.Equal(t, `provider "dns-he-net" {
		username = "user"
		password = "password"
		otp_secret = "otp"
		store_type = "store_type"
	}
	`, c.ProviderConfig("store_type"))

	authObj, err := c.Auth(auth.Simple)
	require.NoError(t, err, "This shouldn't fail")
	assert.Equal(t, "user", authObj.User)
	assert.Equal(t, "password", authObj.Password)
	assert.Equal(t, "otp", authObj.OTPKey.Secret())
}

func TestArpaDomain(t *testing.T) {
	testCases := map[string]bool{
		"0.0.0.0.ip6.arpa":                         true,
		"f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.f.ip6.arpa": true,
		"g.g.g.g.ip6.arpa":                         false,
	}

	for domain, isValid := range testCases {
		err := test_cfg.ValidateArpaDomain(domain)
		if isValid {
			require.NoError(t, err, "expected domain %q to be valid, got error: %v", domain, err)
		} else {
			require.Error(t, err, "expected domain %q to be invalid, but got no error", domain)
		}
	}
}

func TestReverseString(t *testing.T) {
	testCases := map[string]string{
		"abc":   "cba",
		"hello": "olleh",
		"":      "",
	}

	for input, expected := range testCases {
		assert.Equal(t, expected, test_cfg.ReverseString(input), "expected ReverseString(%q) to be %q", input, expected)
	}
}

func TestGenerateArpaSubDomains(t *testing.T) {
	c := test_cfg.ZoneCfg{
		Name: "1.0.0.0.3.1.f.1.0.7.4.0.1.0.0.2.ip6.arpa",
	}

	subdomains := c.RandArpaSubs(16, 16) // Generate 16 subdomains with 1 byte (2 hex chars)
	assert.Len(t, subdomains, 16)
	for _, subdomain := range subdomains {
		assert.NoError(t, test_cfg.ValidateArpaDomain(subdomain), "generated subdomain %q is not a valid ARPA domain", subdomain)
		assert.True(t, strings.HasSuffix(subdomain, c.Name), "generated subdomain %q does not have the correct suffix", subdomain)
	}
}

// func TestPrint(t *testing.T) {
// 	fmt.Println(os.Getenv("DNSHENET_TEST_CONFIG_PATH"))
// 	fmt.Println(test_cfg.Config.DataSources.DomainZones)
// }
