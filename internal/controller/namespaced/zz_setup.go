// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	balancer "github.com/miaits/provider-hetzner/internal/controller/namespaced/loadbalancer/balancer"
	balancernetwork "github.com/miaits/provider-hetzner/internal/controller/namespaced/loadbalancer/balancernetwork"
	balancerservice "github.com/miaits/provider-hetzner/internal/controller/namespaced/loadbalancer/balancerservice"
	balancertarget "github.com/miaits/provider-hetzner/internal/controller/namespaced/loadbalancer/balancertarget"
	network "github.com/miaits/provider-hetzner/internal/controller/namespaced/network/network"
	route "github.com/miaits/provider-hetzner/internal/controller/namespaced/network/route"
	subnet "github.com/miaits/provider-hetzner/internal/controller/namespaced/network/subnet"
	providerconfig "github.com/miaits/provider-hetzner/internal/controller/namespaced/providerconfig"
	group "github.com/miaits/provider-hetzner/internal/controller/namespaced/server/group"
	networkserver "github.com/miaits/provider-hetzner/internal/controller/namespaced/server/network"
	server "github.com/miaits/provider-hetzner/internal/controller/namespaced/server/server"
	snapshot "github.com/miaits/provider-hetzner/internal/controller/namespaced/server/snapshot"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		balancer.Setup,
		balancernetwork.Setup,
		balancerservice.Setup,
		balancertarget.Setup,
		network.Setup,
		route.Setup,
		subnet.Setup,
		providerconfig.Setup,
		group.Setup,
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
		network.SetupGated,
		route.SetupGated,
		subnet.SetupGated,
		providerconfig.SetupGated,
		group.SetupGated,
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
