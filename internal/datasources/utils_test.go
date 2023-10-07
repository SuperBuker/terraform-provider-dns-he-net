package datasources_test

import (
	"strconv"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
)

var (
	// Capitalised variables are accessed by the entire the test package
	_datasources   = test_utils.Config.DataSouces
	ProviderConfig = _datasources.Account.ProviderConfig("simple")
	Account        = _datasources.Account
	Records        = _datasources.Records
	Zone           = _datasources.Zone
	ZonesCount     = _datasources.ZonesCount
)

func toString(i uint) string {
	return strconv.FormatUint(uint64(i), 10)
}
