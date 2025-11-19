// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Package xds provides an implementation of a gRPC service that exports Envoy's
// xDS API for config discovery. Specifically we support the Aggregated
// Discovery Service (ADS) only as we control all config.
//
// A full description of the XDS protocol can be found at
// https://www.envoyproxy.io/docs/envoy/latest/api-docs/xds_protocol
//
// xds.Server also support ext_authz network filter API to authorize incoming
// connections to Envoy.

/*
# Module: agent/xds/xds.go
Xds module for agent layer

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/xds.go> a code:Module ;
    code:name "agent/xds/xds.go" ;
    code:description "Xds module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds
