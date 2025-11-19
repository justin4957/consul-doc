// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/util/net.go
Net module for internal layer

## Linked Modules
- [testing/deployer/util/internal/ipamutils](../testing/deployer/util/internal/ipamutils)

## Tags
internal

## Exports
GetPossibleDockerNetworkSubnets

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/util/net.go> a code:Module ;
    code:name "testing/deployer/util/net.go" ;
    code:description "Net module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "testing/deployer/util/internal/ipamutils" ;
        code:path "../testing/deployer/util/internal/ipamutils"
    ] ;
    code:exports :GetPossibleDockerNetworkSubnets ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package util

import (
	"github.com/hashicorp/consul/testing/deployer/util/internal/ipamutils"
)

// GetPossibleDockerNetworkSubnets returns a copy of the global-scope network list.
func GetPossibleDockerNetworkSubnets() map[string]struct{} {
	list := ipamutils.GetGlobalScopeDefaultNetworks()

	out := make(map[string]struct{})
	for _, ipnet := range list {
		subnet := ipnet.String()
		out[subnet] = struct{}{}
	}
	return out
}
