package resources_test

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
	user := os.Getenv("DNSHENET_USER")
	password := os.Getenv("DNSHENET_PASSWD")
	otp := os.Getenv("DNSHENET_OTP")

	auth, err := client.NewAuth(user, password, otp, client.CookieStore.Simple)

	if err != nil {
		log.Printf("Auth init failed : %s", err.Error())
		os.Exit(1)
		return
	}

	client, err := client.NewClient(context.TODO(), auth,
		logging.NewZerolog(zerolog.DebugLevel, false))

	if err != nil {
		log.Printf("Client init failed : %s", err.Error())
		os.Exit(1)
		return
	}

	// Ensure authentication works
	if _, err = client.GetZones(context.TODO()); err != nil {
		log.Printf("Authentication failed : %s", err.Error())
		os.Exit(1)
		return
	}

	exitVal := m.Run()

	os.Exit(exitVal)
}
