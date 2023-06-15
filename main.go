package main

import (
	"context"
	"flag"
	"log"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Run "go generate" to format example terraform files and generate the docs for the registry/website

// If you do not have terraform installed, you can remove the formatting command, but its suggested to
// ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ./examples/

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	debug   bool
)

func BuildFlags() internal.BuildFlags {
	return internal.BuildFlags{
		Version: version,
		Commit:  commit,
		Date:    date,
	}
}

func RunFlags() internal.RunFlags {
	return internal.RunFlags{
		Debug: debug,
	}
}

func main() {
	provider := internal.New(BuildFlags(), RunFlags())

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/SuperBuker/dns-he-net",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider, opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()
}
