package placementgroup

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for placement group cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_placement_group", func(r *config.Resource) {
		r.ShortGroup = "server"
	})
}
