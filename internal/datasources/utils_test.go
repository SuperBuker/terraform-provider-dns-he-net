package datasources_test

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
)

var (
	// Capitalised variables are accessed by the entire the test package
	_datasources   = test_cfg.Config.DataSouces
	ProviderConfig = _datasources.Account.ProviderConfig("simple")
	Account        = _datasources.Account
	Records        = _datasources.Records
	Zone           = _datasources.Zone
	ZonesCount     = _datasources.ZonesCount
)

func toString(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}
