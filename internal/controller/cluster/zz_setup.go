// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	balancer "github.com/miaits/provider-hetzner/internal/controller/cluster/loadbalancer/balancer"
	balancernetwork "github.com/miaits/provider-hetzner/internal/controller/cluster/loadbalancer/balancernetwork"
	balancerservice "github.com/miaits/provider-hetzner/internal/controller/cluster/loadbalancer/balancerservice"
	balancertarget "github.com/miaits/provider-hetzner/internal/controller/cluster/loadbalancer/balancertarget"
	firewall "github.com/miaits/provider-hetzner/internal/controller/cluster/network/firewall"
	network "github.com/miaits/provider-hetzner/internal/controller/cluster/network/network"
	route "github.com/miaits/provider-hetzner/internal/controller/cluster/network/route"
	subnet "github.com/miaits/provider-hetzner/internal/controller/cluster/network/subnet"
	providerconfig "github.com/miaits/provider-hetzner/internal/controller/cluster/providerconfig"
	group "github.com/miaits/provider-hetzner/internal/controller/cluster/server/group"
	ip "github.com/miaits/provider-hetzner/internal/controller/cluster/server/ip"
	networkserver "github.com/miaits/provider-hetzner/internal/controller/cluster/server/network"
	server "github.com/miaits/provider-hetzner/internal/controller/cluster/server/server"
	snapshot "github.com/miaits/provider-hetzner/internal/controller/cluster/server/snapshot"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		balancer.Setup,
		balancernetwork.Setup,
		balancerservice.Setup,
		balancertarget.Setup,
		firewall.Setup,
		network.Setup,
		route.Setup,
		subnet.Setup,
		providerconfig.Setup,
		group.Setup,
		ip.Setup,
		networkserver.Setup,
		server.Setup,
		snapshot.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		balancer.SetupGated,
		balancernetwork.SetupGated,
		balancerservice.SetupGated,
		balancertarget.SetupGated,
		firewall.SetupGated,
		network.SetupGated,
		route.SetupGated,
		subnet.SetupGated,
		providerconfig.SetupGated,
		group.SetupGated,
		ip.SetupGated,
		networkserver.SetupGated,
		server.SetupGated,
		snapshot.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
