package volume

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for volume namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_volume", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
