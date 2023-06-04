package main

import (
	"context"
	"flag"
	"log"

	"github.com/SuperBuker/terraform-provider-dns-he-net/internal"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

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
