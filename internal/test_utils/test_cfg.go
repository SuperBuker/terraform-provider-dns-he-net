package test_utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	//"log"

	"github.com/SuperBuker/terraform-provider-dns-he-net/client/auth"
	"github.com/sethvargo/go-envconfig"
)

var (
	// Config is a shared configuration to combine with the actual
	// test configuration so the dns.he.net client is properly configured.
	Config TestCfg
)

// AccountCfg contains the configuration for the account to use for testing.
// The account configuration can be loaded from a JSON file or from the
// DHN_ environment variables.
type AccountCfg struct {
	User     string `json:"user" env:"DNSHENET_USER"`
	Password string `json:"password" env:"DNSHENET_PASSWD"`
	OTP      string `json:"otp" env:"DNSHENET_OTP"`
	ID       string `json:"id" env:"DNSHENET_ACCOUNT_ID"`
}

// loadENV loads the account configuration from the environment variables.
func (c *AccountCfg) loadENV() error {
	ctx := context.Background()
	return envconfig.Process(ctx, c)
}

// ProviderConfig is a shared configuration to combine with the actual
// test configuration so the dns.he.net client is properly configured.
func (c AccountCfg) ProviderConfig(store_type string) string {
	return fmt.Sprintf(`provider "dns-he-net" {
		username = %q
		password = %q
		otp_secret = %q
		store_type = %q
	}
	`, c.User, c.Password, c.OTP, store_type)
}

// Auth returns an auth.Auth instance for the account configuration.
func (c AccountCfg) Auth(store_type auth.AuthStore) (auth.Auth, error) {
	return auth.NewAuth(c.User, c.Password, c.OTP, store_type)
}

type RecordCfg struct {
	ID     uint   `json:"id"`
	Domain string `json:"domain"`
	Data   string `json:"data"`
	TTL    uint   `json:"ttl"`
}

type ZoneCfg struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (c ZoneCfg) Sub(subdomain string) string {
	return fmt.Sprintf("%s.%s", subdomain, c.Name)
}

func (c ZoneCfg) RandSub(prefix string, size int, len int) []string {
	return generateSubDomains(fmt.Sprintf("%s.%s", prefix, c.Name), size, len)
}

type DataSoucesTestCfg struct {
	Account    AccountCfg           `json:"account"`
	Zone       ZoneCfg              `json:"zone"`
	ZonesCount uint                 `json:"zones_count"`
	Records    map[string]RecordCfg `json:"records"`
}

type ResourceTestCFG struct {
	Account AccountCfg `json:"account"`
	Zone    ZoneCfg    `json:"zone"`
}

type TestCfg struct {
	DataSouces DataSoucesTestCfg `json:"datasources"`
	Resources  ResourceTestCFG   `json:"resources"`
}

// Load loads the test configuration from the JSON file.
func (c *TestCfg) Load(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		log.Fatal(err)
	}

	// Load the test configuration from the environment variables
	for _, v := range []*AccountCfg{
		&c.DataSouces.Account,
		&c.Resources.Account,
	} {
		err = v.loadENV()
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	// Load the test configuration file
	config_path := os.Getenv("DNSHENET_TEST_CONFIG_PATH")
	if config_path == "" {
		log.Fatal("DNSHENET_TEST_CONFIG_PATH is not set")
	}

	if err := Config.Load(config_path); err != nil {
		log.Fatal(err)
	}
}
