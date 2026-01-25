package sshkey

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for ssh key namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_ssh_key", func(r *config.Resource) {
		r.ShortGroup = "server"
	})
}
