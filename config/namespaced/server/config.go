package server

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for server namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_server", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["firewall_ids"] = config.Reference{
			TerraformName: "hcloud_firewall",
		}
		r.References["placement_group_id"] = config.Reference{
			TerraformName: "hcloud_placement_group",
		}
		r.References["ssh_keys"] = config.Reference{
			TerraformName: "hcloud_ssh_key",
		}
	})
}
