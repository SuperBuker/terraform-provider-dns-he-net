package test_cfg_test

import (
	"os"
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
