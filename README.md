# Terraform Provider: Hurricane Electric DNS

[![Actions Status](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions/workflows/golang.yaml/badge.svg?branch=master)](https://github.com/SuperBuker/terraform-provider-dns-he-net/actions)
[![codecov](https://codecov.io/gh/SuperBuker/terraform-provider-dns-he-net/graph/badge.svg?token=ODPKLRKW5Q)](https://codecov.io/gh/SuperBuker/terraform-provider-dns-he-net)
[![Go Report Card](https://goreportcard.com/badge/github.com/SuperBuker/terraform-provider-dns-he-net)](https://goreportcard.com/report/github.com/SuperBuker/terraform-provider-dns-he-net)
[![GitHub release](https://img.shields.io/github/v/tag/Superbuker/terraform-provider-dns-he-net?label=release)](https://github.com/SuperBuker/terraform-provider-dns-he-net/releases)
![Terraform registry downloads](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fregistry.terraform.io%2Fv2%2Fproviders%2F4236%2Fdownloads%2Fsummary&query=%24.data.attributes.total&logo=terraform&label=downloads%20&color=%237B42BC&link=https%3A%2F%2Fregistry.terraform.io%2Fproviders%2FSuperBuker%2Fdns-he-net%2Flatest)
[![License](https://img.shields.io/github/license/SuperBuker/terraform-provider-dns-he-net.svg)]()

Terraform Provider for setting DNS records in Hurricane Electric.

## Usage

### Terraform Registry

The documentation can be found at [registry.terraform.io](https://registry.terraform.io/providers/SuperBuker/dns-he-net/latest).

Examples can be found in the [examples](./examples) folder.

### Provider Configuration
    
```terraform
terraform {
  required_providers {
    dns-he-net = {
      source = "SuperBuker/dns-he-net"
    }
  }
}

provider "dns-he-net" {
  username   = "username"
  password   = "password"
  otp_secret = "otp_secret" # optional, only if enabled
  store_type = "encrypted"  # optional, default: "encrypted"
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

### Building the Provider

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

Everything related with testing is available in the [testing](./testing) folder.

[Bug reports](https://github.com/SuperBuker/terraform-provider-dns-he-net/issues/new/choose) are welcome. :wink:

### Installing the Provider

After the build is complete, copy the binary into the `~/.terraform.d/plugins` directory

```sh
$ cp ./terraform-provider-dns-he-net ~/.terraform.d/plugins/terraform-provider-dns-he-net
```

## Contribute

If you'd like to thank the developers behind this project leave a [GitHub Star](https://github.com/SuperBuker/terraform-provider-dns-he-net/stargazers) and spread the word.  
This is just a humble hobby project, we are not looking for donations.

## License

[GPLv3](./LICENSE)

## References

- [Hurricane Electric DNS](https://dns.he.net)
- [Terraform Plugin Framework](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework)
