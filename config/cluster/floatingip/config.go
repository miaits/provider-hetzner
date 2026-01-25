package floatingip

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for floating ip cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_floating_ip", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
