package floatingipassignment

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for floating IP assignment cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_floating_ip_assignment", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["floating_ip_id"] = config.Reference{
			TerraformName: "hcloud_floating_ip",
		}
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
