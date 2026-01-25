package uploadedcertificate

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure adds configurations for uploaded certificate namespaced.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_uploaded_certificate", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
