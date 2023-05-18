package internal_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Provider ok
			{
				Config: fmt.Sprintf(`
				provider "dns-he-net" {
				  username = %q
				  password = %q
				  otp_secret = %q
				  store_type = "simple"
				}
				`, os.Getenv("DNSHENET_USER"), os.Getenv("DNSHENET_PASSWD"), os.Getenv("DNSHENET_OTP")),
			},
		},
	})
}

func TestAccProviderStoreErr(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Provider config error
			{
				Config: fmt.Sprintf(`
				provider "dns-he-net" {
				  username = %q
				  password = %q
				  otp_secret = %q
				  store_type = "x"
				}

				data "dns-he-net_account" "test" {}
				`, os.Getenv("DNSHENET_USER"), os.Getenv("DNSHENET_PASSWD"), os.Getenv("DNSHENET_OTP")),

				ExpectError: regexp.MustCompile("Invalid Attribute Value Match"),
			},
		},
	})
}

func TestAccProviderAuthErr(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: test_utils.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Provider auth error
			{
				Config: `
				provider "dns-he-net" {
					username = "username"
					password = "password"
				}
				
				data "dns-he-net_account" "test" {}
				`,

				ExpectError: regexp.MustCompile("Unable to Create dns.he.net API Client"),
			},
		},
	})
}
