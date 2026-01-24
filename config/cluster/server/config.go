package server

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for server cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_server", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["placement_group_id"] = config.Reference{
			TerraformName: "hcloud_placement_group",
		}
	})
}
