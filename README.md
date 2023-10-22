# Terraform Provider: Hurricane Electric DNS

[![Actions Status](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions/workflows/golang.yaml/badge.svg?branch=master)](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions)
[![codecov](https://codecov.io/gh/SuperBuker/terraform-provider-dns-he-net/graph/badge.svg?token=ODPKLRKW5Q)](https://codecov.io/gh/SuperBuker/terraform-provider-dns-he-net)
[![Go Report Card](https://goreportcard.com/badge/github.com/SuperBuker/terraform-provider-dns-he-net)](https://goreportcard.com/report/github.com/SuperBuker/terraform-provider-dns-he-net)
[![GitHub release](https://img.shields.io/github/v/tag/Superbuker/terraform-provider-dns-he-net?label=release)](https://github.com/SuperBuker/terraform-provider-dns-he-net/releases)
[![License](https://img.shields.io/github/license/SuperBuker/terraform-provider-dns-he-net.svg)]()

Terraform Provider for setting DNS records in Hurricane Electric.

## Usage

### Terraform Registry

The documentation can be found at [registry.terraform.io](https://registry.terraform.io/providers/SuperBuker/dns-he-net/latest).

Examples can be found in the [examples](./examples) folder.

### Provider Configuration
    
```hcl
terraform {
  required_providers {
    dns-he-net = {
      source = "SuperBuker/dns-he-net"
    }
  }
}

provider "dns-he-net" {
  username = "username"
  password = "password"
}

resource "dns-he-net_a" "example" {
  zone_id = 123456
  domain  = "example.com"
  ttl     = 86400
  data    = "1.2.3.4"
}
```

Then run `terraform init` to download and install the provider.

## Development

### Requirements

- [Go](https://golang.org/) 1.20 (to build the provider plugin)
- [Terraform](https://www.terraform.io/downloads.html) >= v1.0

### Building The Provider

Clone the repository.

```sh
$ git clone git@github.com:SuperBuker/terraform-provider-dns-he-net.git
```

Enter the provider directory and build the provider

```sh
$ cd ./terraform-provider-dns-he-net
$ go build -o terraform-provider-dns-he-net
```

### Testing

> **Note:** This section is under construction, the testing procedure has been entirely reworked to make the test entironment reproducible.
> A preview is currently available in the [testing](https://github.com/SuperBuker/terraform-provider-dns-he-net/tree/master/testing) folder.

In order to run the full suite of Acceptance tests, the following environment variables must be set:

- `DNSHENET_USER` (mandatory)
- `DNSHENET_PASSWD` (mandatory)
- `DNSHENET_OTP` (optional, only required if the OTP auth is enabled)
- `DNSHENET_ACCOUNT_ID` (mandatory, used for validating the testing results)

```sh
$ TF_ACC=1 go test -v ./...
```

### Installing the Provider

After the build is complete, copy the binary into the `~/.terraform.d/plugins` directory

```sh
$ cp ./terraform-provider-dns-he-net ~/.terraform.d/plugins/terraform-provider-dns-he-net
```

## License

[GPLv3](./LICENSE)

## References

- [Hurricane Electric DNS](https://dns.he.net)
- [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework)
