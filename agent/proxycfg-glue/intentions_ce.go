// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg-glue/intentions_ce.go
Intentions Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)
- [proto/private/pbsubscribe](../proto/private/pbsubscribe)

## Tags
agent, networking, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/intentions_ce.go> a code:Module ;
    code:name "agent/proxycfg-glue/intentions_ce.go" ;
    code:description "Intentions Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pbsubscribe" ;
        code:path "../proto/private/pbsubscribe"
    ] ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbsubscribe"
)

func (s serverIntentions) buildSubjects(serviceName string, entMeta acl.EnterpriseMeta) []*pbsubscribe.NamedSubject {
	// Based on getIntentionPrecedenceMatchServiceNames in the state package.
	if serviceName == structs.WildcardSpecifier {
		return []*pbsubscribe.NamedSubject{
			{
				Key:       structs.WildcardSpecifier,
				Namespace: entMeta.NamespaceOrDefault(),
				Partition: entMeta.PartitionOrDefault(),
				PeerName:  structs.DefaultPeerKeyword,
			},
		}
	}

	return []*pbsubscribe.NamedSubject{
		{
			Key:       serviceName,
			Namespace: entMeta.NamespaceOrDefault(),
			Partition: entMeta.PartitionOrDefault(),
			PeerName:  structs.DefaultPeerKeyword,
		},
		{
			Key:       structs.WildcardSpecifier,
			Namespace: entMeta.NamespaceOrDefault(),
			Partition: entMeta.PartitionOrDefault(),
			PeerName:  structs.DefaultPeerKeyword,
		},
	}
}
