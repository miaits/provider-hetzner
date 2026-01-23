package loadbalancerservice

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for load balancer service cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer_service", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "hcloud_load_balancer",
		}
	})
}
