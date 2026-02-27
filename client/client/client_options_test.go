package client_test

import (
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientOptions(t *testing.T) {
	_resources := test_cfg.Config.Resources
	account := _resources.Account

	authObj, err := account.Auth(auth.Simple)
	require.NoError(t, err)

	options := make(client.Options, 0)

	options = append(options,
		client.WithProxy("http://localhost:8080"), // Example proxy
		client.WithUserAgent("Random User Agent"),
		client.WithDebug(),
	)

	// Here we just want to test that the options don't break the client creation
	cli, err := client.NewClient(t.Context(), authObj, logging.NewZerolog(zerolog.DebugLevel, false), options...)
	require.NoError(t, err)
	assert.NotNil(t, cli)
}
