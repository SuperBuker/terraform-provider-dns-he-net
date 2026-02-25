package test_cfg

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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
	return envconfig.Process(context.Background(), c)
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
	ID        uint              `json:"id"`
	Domain    string            `json:"domain"`
	Data      string            `json:"data"`
	TTL       uint              `json:"ttl"`
	Dynamic   bool              `json:"dynamic"`
	ExtraArgs map[string]string `json:"extra_args,omitempty"`
}

type ZoneCfg struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Records     map[string]RecordCfg `json:"records"`
	RecordCount uint                 `json:"record_count"`
}

func (c ZoneCfg) Sub(subdomain string) string {
	return fmt.Sprintf("%s.%s", subdomain, c.Name)
}

func (c ZoneCfg) RandSubs(prefix string, bound int, count int) []string {
	return generateSubDomains(fmt.Sprintf("%s.%s", prefix, c.Name), bound, count)
}

func (c ZoneCfg) RandArpaSubs(bytes int, count int) []string {
	return generateArpaSubDomains(c.Name, bytes, count)
}

type DataSourcesDomainZoneCfg struct {
	Ok                ZoneCfg `json:"ok"`
	PendingDelegation ZoneCfg `json:"pending_delegation"`
}

type DataSourcesArpaZoneCfg struct {
	Ok ZoneCfg `json:"ok"`
	//PendingDelegation ZoneCfg `json:"pending_delegation"`
}

type NetworkPrefixCfg struct {
	ID      uint   `json:"id"`
	Value   string `json:"value"`
	Enabled bool   `json:"enabled"`
}

type DataSourcesNetworkPrefixCfg struct {
	Ok NetworkPrefixCfg `json:"ok"`
	//TODO: PendingActivation NetworkPrefixCfg `json:"pending_activation"`
}

type DataSourcesTestCfg struct {
	Account              AccountCfg                  `json:"account"`
	DomainZones          DataSourcesDomainZoneCfg    `json:"domain_zones"`
	DomainZonesCount     uint                        `json:"domain_zones_count"`
	NetworkPrefixes      DataSourcesNetworkPrefixCfg `json:"network_prefixes"`
	NetworkPrefixesCount uint                        `json:"network_prefixes_count"`
	ArpaZones            DataSourcesArpaZoneCfg      `json:"arpa_zones"`
	ArpaZonesCount       uint                        `json:"arpa_zones_count"`
}

type ResourceTestCfg struct {
	Account    AccountCfg `json:"account"`
	DomainZone ZoneCfg    `json:"domain_zone"`
	ArpaZone   ZoneCfg    `json:"arpa_zone"`
}

type TestCfg struct {
	DataSources DataSourcesTestCfg `json:"datasources"`
	Resources   ResourceTestCfg    `json:"resources"`
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
		&c.DataSources.Account,
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
