// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/testing_upstreams_ce.go
Testing Upstreams Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)
- [proto/private/pbpeering](../proto/private/pbpeering)

## Tags
agent, networking, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/testing_upstreams_ce.go> a code:Module ;
    code:name "agent/proxycfg/testing_upstreams_ce.go" ;
    code:description "Testing Upstreams Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pbpeering" ;
        code:path "../proto/private/pbpeering"
    ] ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

import (
	"github.com/mitchellh/go-testing-interface"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbpeering"
)

func extraDiscoChainConfig(t testing.T, variation string, entMeta acl.EnterpriseMeta) ([]structs.ConfigEntry, []*pbpeering.Peering) {
	t.Fatalf("unexpected variation: %q", variation)
	return nil, nil
}

func extraUpdateEvents(t testing.T, variation string, dbUID UpstreamID) []UpdateEvent {
	t.Fatalf("unexpected variation: %q", variation)
	return nil
}
