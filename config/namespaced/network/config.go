package network

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for network namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
