package loadbalancernetwork

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for load balancer network cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer_network", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
		r.LateInitializer.IgnoredFields = append(r.LateInitializer.IgnoredFields, "network_id")
		r.References["load_balancer_id"] = config.Reference{
			TerraformName: "hcloud_load_balancer",
		}
		r.References["network_id"] = config.Reference{
			TerraformName: "hcloud_network",
		}
		r.References["subnet_id"] = config.Reference{
			TerraformName: "hcloud_network_subnet",
		}
	})
}
