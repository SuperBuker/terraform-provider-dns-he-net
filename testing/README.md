# How to Test dns.he.net Terraform Provider

## Setup

There are two levels of testing:
- Client API unit testing
- Terraform provider integration testing

### Client API Unit Testing

The client needs valid credentials to authenticate against the dns.he.net api.  
The credentials are stored in a file called `./testing/files/test_config.json`,
an example is available in [`./testing/files/test_config_example.json`](./files/test_config_example.json).

The client makes exclusive use of the datasource credentials to authenticate against dns.he.net.

It's also possible to set the credentials via environment variables:

```sh
$ export DNSHENET_USER='username'
$ export DNSHENET_PASSWD='password'
$ export DNSHENET_OTP='opt_secret'             # optional, only required if enabled in the account
$ export DNSHENET_ACCOUNT_ID='account_id'
$ export DNSHENET_TEST_CONFIG_PATH='path'      # optional, "testing/files/test_config.json" by default
```

### Terraform Provider Integration Testing

The Terraform provider integration testing is done via the [Terraform Plugin Testing Framework](https://developer.hashicorp.com/terraform/plugin/testing).  
The testing performs real interactions against the service API, so it not only
requires valid credentials to authenticate against dns.he.net, but also an
isolated environment to avoid any side effects.

Additionally, two domains are required to perform the integration testing, both
registered on dns.he.net:  
The first domain is used for testing the datasources, thus the records must exist and be valid.  
The second domain is used for testing the resources, thus no records must exist.

This repository contains the terraform manifests to setup the testing environment.  
The manifests are located in [`./testing/infrastructure`](./infrastructure).

To provision the testing environment, following these steps:

**1. Create the file `./testing/infrastructure/terraform.tfvars` and fill it with the following content:**

```terraform
account = {
    "username" : "username",
    "password" : "password",
    "otp_secret" : "otp_secret",               # optional, only required if enabled in the account
    "store_type" : "simple",                   # optional, default "encrypted"
    "mask_creds" : "false"                     # optional, default "true", to mask the credentials in the config file
}

datasources_zone = "example.com"               # domain used for testing the datasources
resources_zone   = "example.org"               # domain used for testing the resources
config_file      = "../files/test_config.json" # optional, "../files/test_config.json" by default
```

**2. Run the following commands:**
    
```sh
$ cd ./testing/infrastructure
$ terraform init
$ terraform apply
```

**3. Setup the env vars for the integration testing:**

All these vars can either be provisioned though the `test_config.json` file or via environment variables.

```sh
$ export DNSHENET_USER='username'
$ export DNSHENET_PASSWD='password'
$ export DNSHENET_OTP='opt_secret'             # optional, only required if enabled in the account
$ export DNSHENET_ACCOUNT_ID='account_id'
```

Aditionally, the following env vars are required for the integration testing:

```sh
$ export TF_ACC='1'                            # required for activating the integration testing
$ export DNSHENET_TEST_CONFIG_PATH='path'      # optional, "testing/files/test_config.json" by default
```

## Testing

Once the config files and environment variables are setup, the testing can be performed as any other golang project:

```sh
$ go test -v ./...
```

> **Happy bug hunting!**
