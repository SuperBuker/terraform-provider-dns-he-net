<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# test\_utils

```go
import "github.com/SuperBuker/terraform-provider-dns-he-net/internal/test_utils"
```

## Index

- [Variables](<#variables>)


## Variables

```go
var (
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

    TestAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)
)
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)