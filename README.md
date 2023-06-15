# Terraform Provider: Hurricane Electric DNS

[![Actions Status](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions/workflows/golang.yaml/badge.svg?branch=master)](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions)
[![GitHub release](https://img.shields.io/github/v/tag/Superbuker/terraform-provider-dns-he-net?label=release)](https://github.com/SuperBuker/terraform-provider-dns-he-net/releases)
[![license](https://img.shields.io/github/license/SuperBuker/terraform-provider-dns-he-net.svg)]()
[![Go Report Card](https://goreportcard.com/badge/github.com/SuperBuker/terraform-provider-dns-he-net)](https://goreportcard.com/report/github.com/SuperBuker/terraform-provider-dns-he-net)

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

In order to run the full suite of Acceptance tests, the following environment variables must be set:

- `DNSHENET_USER`
- `DNSHENET_PASSWD`
- `DNSHENET_OTP`

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
