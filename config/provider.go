package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	loadBalancerCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancer"
	loadBalancerNetworkCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancernetwork"
	loadBalancerServiceCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancerservice"
	loadBalancerTargetCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancertarget"
	networkCluster "github.com/miaits/provider-hetzner/config/cluster/network"
	networkRouteCluster "github.com/miaits/provider-hetzner/config/cluster/networkroute"
	networkSubnetCluster "github.com/miaits/provider-hetzner/config/cluster/networksubnet"
	serverNetworkCluster "github.com/miaits/provider-hetzner/config/cluster/servernetwork"
	serverCluster "github.com/miaits/provider-hetzner/config/cluster/server"
	loadBalancerNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancer"
	loadBalancerNetworkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancernetwork"
	loadBalancerServiceNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancerservice"
	loadBalancerTargetNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancertarget"
	networkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/network"
	networkRouteNamespaced "github.com/miaits/provider-hetzner/config/namespaced/networkroute"
	networkSubnetNamespaced "github.com/miaits/provider-hetzner/config/namespaced/networksubnet"
	serverNetworkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/servernetwork"
	serverNamespaced "github.com/miaits/provider-hetzner/config/namespaced/server"
)

const (
	resourcePrefix = "hetzner"
	modulePath     = "github.com/miaits/provider-hetzner"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("hetzner.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		networkCluster.Configure,
		networkSubnetCluster.Configure,
		networkRouteCluster.Configure,
		loadBalancerCluster.Configure,
		loadBalancerNetworkCluster.Configure,
		loadBalancerServiceCluster.Configure,
		loadBalancerTargetCluster.Configure,
		serverCluster.Configure,
		serverNetworkCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("hetzner.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		networkNamespaced.Configure,
		networkSubnetNamespaced.Configure,
		networkRouteNamespaced.Configure,
		loadBalancerNamespaced.Configure,
		loadBalancerNetworkNamespaced.Configure,
		loadBalancerServiceNamespaced.Configure,
		loadBalancerTargetNamespaced.Configure,
		serverNamespaced.Configure,
		serverNetworkNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
