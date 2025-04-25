package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/code-of-kpp/terraform-provider-playground/internal/provider"
)

// for goreleaser
var version string = "dev"

func main() {
	var debug bool

	flag.BoolVar(
		&debug, "debug", false,
		"Set to true to run the provider with support for debuggers",
	)
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.opentofu.org/example/playground",
		Debug:   debug,
	}
	err := providerserver.Serve(
		context.Background(),
		provider.New(version),
		opts,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
