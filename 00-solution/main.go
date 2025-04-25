package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/code-of-kpp/terraform-provider-playground/internal/provider"
)

// for goreleaser
var version string = "dev"

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.opentofu.org/example/playground",
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
