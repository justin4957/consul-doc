// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/connect.go
Connect API client for service mesh and secure service-to-service communication.

## Linked Modules
- [api/api](./api.go) - Main API client
- [connect/service](../connect/service.go) - Service mesh implementation

## Tags
api, client, connect, service-mesh, tls, mTLS

## Exports
Connect, TelemetryCollectorName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/connect.go> a code:Module ;
    code:name "api/connect.go" ;
    code:description "Connect API client for service mesh and secure service-to-service communication" ;
    code:layer "api" ;
    code:linksTo [
        code:name "api/api" ;
        code:path "./api.go" ;
        code:relationship "Main API client"
    ], [
        code:name "connect/service" ;
        code:path "../connect/service.go" ;
        code:relationship "Service mesh implementation"
    ] ;
    code:exports :Connect, :TelemetryCollectorName ;
    code:tags "api", "client", "connect", "service-mesh", "tls", "mTLS" .
<!-- End LinkedDoc RDF -->
*/
package api

// TelemetryCollectorName is the service name for the Consul Telemetry Collector
const TelemetryCollectorName string = "consul-telemetry-collector"

// Connect can be used to work with endpoints related to Connect, the
// feature for securely connecting services within Consul.
type Connect struct {
	c *Client
}

// Connect returns a handle to the connect-related endpoints
func (c *Client) Connect() *Connect {
	return &Connect{c}
}
