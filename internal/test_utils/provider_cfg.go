package test_utils

import (
	"github.com/SuperBuker/terraform-provider-dns-he-net/internal"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testProvider                    provider.Provider
	TestAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
)

func init() {
	testProvider = internal.New(
		internal.BuildFlags{
			Version: "v0.0.1",
		},
		internal.RunFlags{
			Debug: false,
		})()

	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"dns-he-net": providerserver.NewProtocol6WithError(testProvider),
	}
}
