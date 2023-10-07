package datasources_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client"
	"github.com/SuperBuker/terraform-provider-dns-he-net/client/logging"
	"github.com/rs/zerolog"
)

func TestMain(m *testing.M) {
	auth, err := client.NewAuth(Account.User, Account.Password, Account.OTP, client.CookieStore.Simple)

	if err != nil {
		log.Printf("Auth init failed : %s", err.Error())
		os.Exit(1)
		return
	}

	ctx := context.Background()
	client, err := client.NewClient(ctx, auth,
		logging.NewZerolog(zerolog.DebugLevel, false))

	if err != nil {
		log.Printf("Client init failed : %s", err.Error())
		os.Exit(1)
		return
	}

	// Ensure authentication works
	if _, err = client.GetZones(ctx); err != nil {
		log.Printf("Authentication failed : %s", err.Error())
		os.Exit(1)
		return
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}
