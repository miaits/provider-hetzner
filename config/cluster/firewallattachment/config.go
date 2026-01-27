package firewallattachment

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for firewall attachment cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_firewall_attachment", func(r *config.Resource) {
		r.Kind = "FirewallAttachment"
		r.ShortGroup = "server"
		r.References["firewall_id"] = config.Reference{
			TerraformName: "hcloud_firewall",
		}
		r.References["server_ids"] = config.Reference{
			TerraformName: "hcloud_server",
		}
	})
}
