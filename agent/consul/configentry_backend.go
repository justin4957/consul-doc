// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/configentry_backend.go
Configentry Backend module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/grpc-external/services/configentry](../agent/grpc-external/services/configentry)

## Tags
agent, configuration

## Exports
ConfigEntryBackend, NewConfigEntryBackend

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/configentry_backend.go> a code:Module ;
    code:name "agent/consul/configentry_backend.go" ;
    code:description "Configentry Backend module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/grpc-external/services/configentry" ;
        code:path "../agent/grpc-external/services/configentry"
    ] ;
    code:exports :ConfigEntryBackend, :NewConfigEntryBackend ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/grpc-external/services/configentry"
)

type ConfigEntryBackend struct {
	srv *Server
}

var _ configentry.Backend = (*ConfigEntryBackend)(nil)

// NewConfigEntryBackend returns a configentry.Backend implementation that is bound to the given server.
func NewConfigEntryBackend(srv *Server) *ConfigEntryBackend {
	return &ConfigEntryBackend{
		srv: srv,
	}
}

func (b *ConfigEntryBackend) EnterpriseCheckPartitions(partition string) error {
	return b.enterpriseCheckPartitions(partition)
}

func (b *ConfigEntryBackend) ResolveTokenAndDefaultMeta(token string, entMeta *acl.EnterpriseMeta, authzCtx *acl.AuthorizerContext) (resolver.Result, error) {
	return b.srv.ResolveTokenAndDefaultMeta(token, entMeta, authzCtx)
}
