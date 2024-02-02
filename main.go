package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/bendbennett/terraform-provider-playground/internal/provider"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/bendbennett/playground",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), provider.New(), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
