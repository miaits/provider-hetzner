package firewall

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for firewall cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_firewall", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
