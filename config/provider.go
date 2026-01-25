package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	firewallCluster "github.com/miaits/provider-hetzner/config/cluster/firewall"
	firewallAttachmentCluster "github.com/miaits/provider-hetzner/config/cluster/firewallattachment"
	floatingIPCluster "github.com/miaits/provider-hetzner/config/cluster/floatingip"
	floatingIPAssignmentCluster "github.com/miaits/provider-hetzner/config/cluster/floatingipassignment"
	loadBalancerCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancer"
	loadBalancerNetworkCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancernetwork"
	loadBalancerServiceCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancerservice"
	loadBalancerTargetCluster "github.com/miaits/provider-hetzner/config/cluster/loadbalancertarget"
	managedCertificateCluster "github.com/miaits/provider-hetzner/config/cluster/managedcertificate"
	uploadedCertificateCluster "github.com/miaits/provider-hetzner/config/cluster/uploadedcertificate"
	networkCluster "github.com/miaits/provider-hetzner/config/cluster/network"
	networkRouteCluster "github.com/miaits/provider-hetzner/config/cluster/networkroute"
	networkSubnetCluster "github.com/miaits/provider-hetzner/config/cluster/networksubnet"
	placementGroupCluster "github.com/miaits/provider-hetzner/config/cluster/placementgroup"
	serverCluster "github.com/miaits/provider-hetzner/config/cluster/server"
	serverNetworkCluster "github.com/miaits/provider-hetzner/config/cluster/servernetwork"
	snapshotCluster "github.com/miaits/provider-hetzner/config/cluster/snapshot"
	firewallNamespaced "github.com/miaits/provider-hetzner/config/namespaced/firewall"
	firewallAttachmentNamespaced "github.com/miaits/provider-hetzner/config/namespaced/firewallattachment"
	floatingIPNamespaced "github.com/miaits/provider-hetzner/config/namespaced/floatingip"
	floatingIPAssignmentNamespaced "github.com/miaits/provider-hetzner/config/namespaced/floatingipassignment"
	loadBalancerNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancer"
	loadBalancerNetworkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancernetwork"
	loadBalancerServiceNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancerservice"
	loadBalancerTargetNamespaced "github.com/miaits/provider-hetzner/config/namespaced/loadbalancertarget"
	managedCertificateNamespaced "github.com/miaits/provider-hetzner/config/namespaced/managedcertificate"
	uploadedCertificateNamespaced "github.com/miaits/provider-hetzner/config/namespaced/uploadedcertificate"
	networkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/network"
	networkRouteNamespaced "github.com/miaits/provider-hetzner/config/namespaced/networkroute"
	networkSubnetNamespaced "github.com/miaits/provider-hetzner/config/namespaced/networksubnet"
	placementGroupNamespaced "github.com/miaits/provider-hetzner/config/namespaced/placementgroup"
	serverNamespaced "github.com/miaits/provider-hetzner/config/namespaced/server"
	serverNetworkNamespaced "github.com/miaits/provider-hetzner/config/namespaced/servernetwork"
	snapshotNamespaced "github.com/miaits/provider-hetzner/config/namespaced/snapshot"
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
		managedCertificateCluster.Configure,
		uploadedCertificateCluster.Configure,
		firewallCluster.Configure,
		firewallAttachmentCluster.Configure,
		floatingIPCluster.Configure,
		floatingIPAssignmentCluster.Configure,
		placementGroupCluster.Configure,
		snapshotCluster.Configure,
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
		managedCertificateNamespaced.Configure,
		uploadedCertificateNamespaced.Configure,
		firewallNamespaced.Configure,
		firewallAttachmentNamespaced.Configure,
		floatingIPNamespaced.Configure,
		floatingIPAssignmentNamespaced.Configure,
		placementGroupNamespaced.Configure,
		snapshotNamespaced.Configure,
		serverNamespaced.Configure,
		serverNetworkNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
