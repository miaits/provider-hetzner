package servernetwork

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for server network namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_server_network", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "network_id")
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
		r.References["network_id"] = config.Reference{
			TerraformName: "hcloud_network",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "hcloud_network_subnet",
		}
	})
}
