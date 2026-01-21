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
		loadBalancerID, ok := parameters["load_balancer_id"]
		if !ok || loadBalancerID == nil || loadBalancerID == "" {
			return "", fmt.Errorf("load_balancer_id not set")
		}

		networkID := parameters["network_id"]
		if networkID == nil || networkID == "" {
			subnetID, ok := parameters["subnet_id"]
			if !ok || subnetID == nil || subnetID == "" {
				return "", fmt.Errorf("network_id or subnet_id must be set")
			}
			derivedNetworkID, err := networkIDFromSubnetID(subnetID)
			if err != nil {
				return "", err
			}
			networkID = derivedNetworkID
		}

		idParams := make(map[string]any, len(parameters)+2)
		for key, value := range parameters {
			idParams[key] = value
		}
		idParams["load_balancer_id"] = numericIDToString(loadBalancerID)
		idParams["network_id"] = numericIDToString(networkID)

		return fn(ctx, externalName, idParams, terraformProviderConfig)
	}))
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
	case string:
		return v
	case fmt.Stringer:
		return v.String()
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
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
