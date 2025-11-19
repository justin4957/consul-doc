// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !linux

/*
# Module: agent/xds/platform/net_fallback.go
Net Fallback module for agent layer

## Tags
agent, envoy, service-mesh

## Exports
SupportsIPv6

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/platform/net_fallback.go> a code:Module ;
    code:name "agent/xds/platform/net_fallback.go" ;
    code:description "Net Fallback module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :SupportsIPv6 ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package platform

func SupportsIPv6() (bool, error) {
	return true, nil
}
