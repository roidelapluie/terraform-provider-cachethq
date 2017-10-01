package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/roidelapluie/terraform-provider-cachethq/cachethq"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: cachethq.Provider})
}
