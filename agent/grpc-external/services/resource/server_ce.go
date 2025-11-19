// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/grpc-external/services/resource/server_ce.go
Server Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
agent, api, communication, grpc, networking

## Exports
Config

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/resource/server_ce.go> a code:Module ;
    code:name "agent/grpc-external/services/resource/server_ce.go" ;
    code:description "Server Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :Config ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func v2TenancyToV1EntMeta(tenancy *pbresource.Tenancy) *acl.EnterpriseMeta {
	return acl.DefaultEnterpriseMeta()
}

func v1EntMetaToV2Tenancy(reg *resource.Registration, entMeta *acl.EnterpriseMeta, tenancy *pbresource.Tenancy) {
	if (reg.Scope == resource.ScopeNamespace || reg.Scope == resource.ScopePartition) && tenancy.Partition == "" {
		tenancy.Partition = entMeta.PartitionOrDefault()
	}

	if reg.Scope == resource.ScopeNamespace && tenancy.Namespace == "" {
		tenancy.Namespace = entMeta.NamespaceOrDefault()
	}
}

type Config struct {
	Logger   hclog.Logger
	Registry Registry

	// Backend is the storage backend that will be used for resource persistence.
	Backend     Backend
	ACLResolver ACLResolver
	// TenancyBridge temporarily allows us to use V1 implementations of
	// partitions and namespaces until V2 implementations are available.
	TenancyBridge TenancyBridge
}

// FeatureCheck does not apply to the community edition.
func (s *Server) FeatureCheck(reg *resource.Registration) error {
	return nil
}
