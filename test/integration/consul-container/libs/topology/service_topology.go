// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/topology/service_topology.go
Service Topology module for internal layer

## Linked Modules
- [api](../api)
- [test/integration/consul-container/libs/assert](../test/integration/consul-container/libs/assert)
- [test/integration/consul-container/libs/cluster](../test/integration/consul-container/libs/cluster)
- [test/integration/consul-container/libs/service](../test/integration/consul-container/libs/service)

## Tags
internal

## Exports
CreateServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/topology/service_topology.go> a code:Module ;
    code:name "test/integration/consul-container/libs/topology/service_topology.go" ;
    code:description "Service Topology module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ], [
        code:name "test/integration/consul-container/libs/assert" ;
        code:path "../test/integration/consul-container/libs/assert"
    ], [
        code:name "test/integration/consul-container/libs/cluster" ;
        code:path "../test/integration/consul-container/libs/cluster"
    ], [
        code:name "test/integration/consul-container/libs/service" ;
        code:path "../test/integration/consul-container/libs/service"
    ] ;
    code:exports :CreateServices ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package topology

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/api"
	libassert "github.com/hashicorp/consul/test/integration/consul-container/libs/assert"
	libcluster "github.com/hashicorp/consul/test/integration/consul-container/libs/cluster"
	libservice "github.com/hashicorp/consul/test/integration/consul-container/libs/service"
)

// CreateServices
func CreateServices(t *testing.T, cluster *libcluster.Cluster, protocol string) (libservice.Service, libservice.Service) {
	node := cluster.Agents[0]
	client := node.GetClient()

	// Register service as HTTP
	serviceDefault := &api.ServiceConfigEntry{
		Kind:     api.ServiceDefaults,
		Name:     libservice.StaticServerServiceName,
		Protocol: protocol,
	}

	ok, _, err := client.ConfigEntries().Set(serviceDefault, nil)
	require.NoError(t, err, "error writing HTTP service-default")
	require.True(t, ok, "did not write HTTP service-default")

	// Create a service and proxy instance
	serviceOpts := &libservice.ServiceOpts{
		Name:     libservice.StaticServerServiceName,
		ID:       "static-server",
		HTTPPort: 8080,
		GRPCPort: 8079,
	}

	if protocol == "grpc" {
		serviceOpts.RegisterGRPC = true
	}

	// Create a service and proxy instance
	_, serverConnectProxy, err := libservice.CreateAndRegisterStaticServerAndSidecar(node, serviceOpts)
	require.NoError(t, err)

	libassert.CatalogServiceExists(t, client, fmt.Sprintf("%s-sidecar-proxy", libservice.StaticServerServiceName), nil)
	libassert.CatalogServiceExists(t, client, libservice.StaticServerServiceName, nil)

	// Create a client proxy instance with the server as an upstream
	clientConnectProxy, err := libservice.CreateAndRegisterStaticClientSidecar(node, "", false, false, nil)
	require.NoError(t, err)

	libassert.CatalogServiceExists(t, client, fmt.Sprintf("%s-sidecar-proxy", libservice.StaticClientServiceName), nil)

	return serverConnectProxy, clientConnectProxy
}
