package networksubnet

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for network subnet cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network_id"] = config.Reference{
			TerraformName: "hcloud_network",
		}
	})
}
