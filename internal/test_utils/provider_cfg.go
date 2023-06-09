package test_utils

import (
	"fmt"
	"os"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// ProviderConfig is a shared configuration to combine with the actual
	// test configuration so the dns.he.net client is properly configured.
	// It is also possible to use the DHN_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	ProviderConfig = fmt.Sprintf(`provider "dns-he-net" {
		username = %q
		password = %q
		otp_secret = %q
		store_type = "simple"
	}
	`, os.Getenv("DNSHENET_USER"), os.Getenv("DNSHENET_PASSWD"), os.Getenv("DNSHENET_OTP"))

	testProvider                    provider.Provider
	TestAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
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
