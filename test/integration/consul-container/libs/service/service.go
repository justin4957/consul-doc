// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/service/service.go
Service module for internal layer

## Linked Modules
- [api](../api)

## Tags
internal

## Exports
Service

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/service/service.go> a code:Module ;
    code:name "test/integration/consul-container/libs/service/service.go" ;
    code:description "Service module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :Service ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package service

import (
	"context"

	"github.com/hashicorp/consul/api"
)

// Service represents a process that will be registered with the
// Consul catalog, including Consul components such as sidecars and gateways
type Service interface {
	Exec(ctx context.Context, cmd []string) (string, error)
	// Export a service to the peering cluster
	Export(partition, peer string, client *api.Client) error
	GetAddr() (string, int)
	GetAddrs() (string, []int)
	GetPort(port int) (int, error)
	// GetAdminAddr returns the external admin address
	GetAdminAddr() (string, int)
	GetLogs() (string, error)
	GetName() string
	GetServiceName() string
	Start() (err error)
	Stop() (err error)
	Terminate() error
	Restart() error
	GetStatus() (string, error)
}
