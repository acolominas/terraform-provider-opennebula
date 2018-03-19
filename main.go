package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/acolominas/terraform-provider-opennebula/opennebula"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: opennebula.Provider,
	})
}
