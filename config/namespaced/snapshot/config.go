package snapshot

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for snapshot namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_snapshot", func(r *config.Resource) {
		r.ShortGroup = "server"
		r.References["server_id"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
