package loadbalancer

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for network cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_load_balancer", func(r *config.Resource) {
		r.ShortGroup = "loadbalancer"
	})
}
