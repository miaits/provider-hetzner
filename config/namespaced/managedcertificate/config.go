package managedcertificate

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for managed certificate namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_managed_certificate", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
