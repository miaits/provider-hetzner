package config

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"null_resource":                idWithStub(),
	"hcloud_network":               config.IdentifierFromProvider,
	"hcloud_network_subnet":        config.IdentifierFromProvider,
	"hcloud_network_route":         config.IdentifierFromProvider,
	"hcloud_load_balancer":         config.IdentifierFromProvider,
	"hcloud_load_balancer_network": loadBalancerNetworkExternalName(),
	"hcloud_load_balancer_service": loadBalancerServiceExternalName(),
	"hcloud_server":                config.IdentifierFromProvider,
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

func loadBalancerNetworkExternalName() config.ExternalName {
	base := config.TemplatedStringAsIdentifier("", "{{ .parameters.load_balancer_id }}-{{ .parameters.network_id }}")
	return config.NewExternalNameFrom(base, config.WithGetIDFn(func(fn config.GetIDFn, ctx context.Context, externalName string, parameters map[string]any, terraformProviderConfig map[string]any) (string, error) {
		idParams, err := loadBalancerNetworkIDParams(parameters)
		if err != nil {
			return "", err
		}
		return fn(ctx, externalName, idParams, terraformProviderConfig)
	}))
}

func loadBalancerServiceExternalName() config.ExternalName {
	base := config.TemplatedStringAsIdentifier("", "{{ .parameters.load_balancer_id }}__{{ .parameters.listen_port }}")
	return config.NewExternalNameFrom(base, config.WithGetIDFn(func(fn config.GetIDFn, ctx context.Context, externalName string, parameters map[string]any, terraformProviderConfig map[string]any) (string, error) {
		idParams, err := loadBalancerServiceIDParams(parameters)
		if err != nil {
			return "", err
		}
		return fn(ctx, externalName, idParams, terraformProviderConfig)
	}))
}

func loadBalancerNetworkIDParams(parameters map[string]any) (map[string]any, error) {
	loadBalancerID, ok := parameters["load_balancer_id"]
	if !ok || loadBalancerID == nil || loadBalancerID == "" {
		return nil, fmt.Errorf("load_balancer_id not set")
	}

	networkID, err := resolveNetworkID(parameters)
	if err != nil {
		return nil, err
	}

	idParams := make(map[string]any, len(parameters)+2)
	for key, value := range parameters {
		idParams[key] = value
	}
	idParams["load_balancer_id"] = numericIDToString(loadBalancerID)
	idParams["network_id"] = numericIDToString(networkID)

	return idParams, nil
}

func loadBalancerServiceIDParams(parameters map[string]any) (map[string]any, error) {
	loadBalancerID, ok := parameters["load_balancer_id"]
	if !ok || loadBalancerID == nil || loadBalancerID == "" {
		return nil, fmt.Errorf("load_balancer_id not set")
	}

	listenPort, err := resolveListenPort(parameters)
	if err != nil {
		return nil, err
	}

	idParams := make(map[string]any, len(parameters)+2)
	for key, value := range parameters {
		idParams[key] = value
	}
	idParams["load_balancer_id"] = numericIDToString(loadBalancerID)
	idParams["listen_port"] = numericIDToString(listenPort)

	return idParams, nil
}

func resolveListenPort(parameters map[string]any) (any, error) {
	listenPort := parameters["listen_port"]
	if listenPort != nil && listenPort != "" {
		return listenPort, nil
	}

	protocol, ok := parameters["protocol"]
	if !ok || protocol == nil || protocol == "" {
		return nil, fmt.Errorf("listen_port not set and protocol not set")
	}

	switch strings.ToLower(fmt.Sprint(protocol)) {
	case "http":
		return 80, nil
	case "https":
		return 443, nil
	case "tcp":
		return nil, fmt.Errorf("listen_port not set")
	default:
		return nil, fmt.Errorf("listen_port not set and unexpected protocol %q", protocol)
	}
}

func resolveNetworkID(parameters map[string]any) (any, error) {
	networkID := parameters["network_id"]
	if networkID != nil && networkID != "" {
		return networkID, nil
	}

	subnetID, ok := parameters["subnet_id"]
	if !ok || subnetID == nil || subnetID == "" {
		return nil, fmt.Errorf("network_id or subnet_id must be set")
	}

	return networkIDFromSubnetID(subnetID)
}

func networkIDFromSubnetID(subnetID any) (string, error) {
	raw := numericIDToString(subnetID)
	parts := strings.SplitN(raw, "-", 2)
	if len(parts) < 2 || parts[0] == "" {
		return "", fmt.Errorf("unexpected subnet_id format %q", raw)
	}
	return parts[0], nil
}

func numericIDToString(value any) string {
	switch v := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 0, 64)
	case float64:
		return strconv.FormatFloat(v, 'f', 0, 64)
	default:
		return fmt.Sprint(value)
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
