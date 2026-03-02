package datasources_test

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/test_cfg"
)

var (
	// Capitalised variables are accessed by the entire the test package
	_datasources         = test_cfg.Config.DataSources
	ProviderConfig       = _datasources.Account.ProviderConfig("simple")
	Account              = _datasources.Account
	DomainZone           = _datasources.DomainZones.Ok
	DomainZoneRecords    = _datasources.DomainZones.Ok.Records
	domainZonesCount     = _datasources.DomainZonesCount
	ArpaZone             = _datasources.ArpaZones.Ok
	ArpaZoneRecords      = _datasources.ArpaZones.Ok.Records
	arpaZonesCount       = _datasources.ArpaZonesCount
	NetworkPrefix        = _datasources.NetworkPrefixes.Ok
	networkPrefixesCount = _datasources.NetworkPrefixesCount
)
