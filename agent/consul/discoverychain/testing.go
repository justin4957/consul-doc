// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/discoverychain/testing.go
Testing module for agent layer

## Linked Modules
- [agent/configentry](../agent/configentry)
- [agent/structs](../agent/structs)

## Tags
agent, discovery, dns

## Exports
TestCompileConfigEntries

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/discoverychain/testing.go> a code:Module ;
    code:name "agent/consul/discoverychain/testing.go" ;
    code:description "Testing module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/configentry" ;
        code:path "../agent/configentry"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :TestCompileConfigEntries ;
    code:tags "agent", "discovery", "dns" .
<!-- End LinkedDoc RDF -->
*/
package discoverychain

import (
	"github.com/mitchellh/go-testing-interface"
	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/structs"
)

func TestCompileConfigEntries(t testing.T,
	serviceName string,
	evaluateInNamespace string,
	evaluateInPartition string,
	evaluateInDatacenter string,
	evaluateInTrustDomain string,
	setup func(req *CompileRequest),
	set *configentry.DiscoveryChainSet) *structs.CompiledDiscoveryChain {
	if set == nil {
		set = configentry.NewDiscoveryChainSet()
	}
	req := CompileRequest{
		ServiceName:           serviceName,
		EvaluateInNamespace:   evaluateInNamespace,
		EvaluateInPartition:   evaluateInPartition,
		EvaluateInDatacenter:  evaluateInDatacenter,
		EvaluateInTrustDomain: evaluateInTrustDomain,
		Entries:               set,
	}
	if setup != nil {
		setup(&req)
	}

	chain, err := Compile(req)
	require.NoError(t, err)
	return chain
}
