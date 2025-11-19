// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/connect/uri_service_ce.go
Uri Service Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_service_ce.go> a code:Module ;
    code:name "agent/connect/uri_service_ce.go" ;
    code:description "Uri Service Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"strings"

	"github.com/hashicorp/consul/acl"
)

// GetEnterpriseMeta will synthesize an EnterpriseMeta struct from the SpiffeIDService.
// in CE this just returns an empty (but never nil) struct pointer
func (id SpiffeIDService) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return &acl.EnterpriseMeta{}
}

// PartitionOrDefault breaks from CE's pattern of returning empty strings.
// Although CE has no support for partitions, it still needs to be able to
// handle exportedPartition from peered Consul Enterprise clusters in order
// to generate the correct SpiffeID.
func (id SpiffeIDService) PartitionOrDefault() string {
	if id.Partition == "" {
		return "default"
	}

	return strings.ToLower(id.Partition)
}
