package test_utils

import (
	"fmt"
	"os"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var ( // WIP... heavily
	// ProviderConfig is a shared configuration to combine with the actual
	// test configuration so the dns.he.net client is properly configured.
	// It is also possible to use the DHN_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	ProviderConfig = fmt.Sprintf(`
provider "dns-he-net" {
  username = "%s"
  password = "%s"
  otp_secret = "%s"
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
	testProvider = internal.New()
	//testProvider.Configure(context.Background(), provider.ConfigureRequest{}, &provider.ConfigureResponse{})
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"dns-he-net": providerserver.NewProtocol6WithError(testProvider),
	}
}

// TODO: https://github.com/dell/terraform-provider-powermax/blob/db425aa781b36a9f29fba1c65a3f10bba1c6bbbf/powermax/provider_test.go
