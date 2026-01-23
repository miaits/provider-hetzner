package loadbalancertarget

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for load balancer target namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer_target", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "hcloud_load_balancer",
		}
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
