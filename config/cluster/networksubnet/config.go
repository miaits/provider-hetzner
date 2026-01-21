package networksubnet

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for network subnet cluster.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "network"
		r.References["network"] = config.Reference{
			Type: "github.com/miaits/provider-hetzner/apis/namespaced/network/v1alpha1.Network",
		}
	})
}
