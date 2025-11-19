// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/catalog.go
Catalog module for agent layer

## Linked Modules
- [acl](../acl)
- [api](../api)
- [types](../types)

## Tags
agent, data-model, discovery, registry, types

## Exports
AssignServiceManualVIPsRequest, AssignServiceManualVIPsResponse, CatalogContents, CatalogSummary, HealthSummary

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/catalog.go> a code:Module ;
    code:name "agent/structs/catalog.go" ;
    code:description "Catalog module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "api" ;
        code:path "../api"
    ], [
        code:name "types" ;
        code:path "../types"
    ] ;
    code:exports :AssignServiceManualVIPsRequest, :AssignServiceManualVIPsResponse, :CatalogContents, :CatalogSummary, :HealthSummary ;
    code:tags "agent", "data-model", "discovery", "registry", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/types"
)

// These are used to manage the built-in "serfHealth" check that's attached
// to every node in the catalog.
const (
	SerfCheckID           types.CheckID = "serfHealth"
	SerfCheckName                       = "Serf Health Status"
	SerfCheckAliveOutput                = "Agent alive and reachable"
	SerfCheckFailedOutput               = "Agent not live or unreachable"
)

const (
	// These are used to manage the "consul" service that's attached to every
	// Consul server node in the catalog.
	ConsulServiceID   = "consul"
	ConsulServiceName = "consul"
)

type CatalogContents struct {
	Nodes    []*Node
	Services []*ServiceNode
	Checks   []*HealthCheck
}

type CatalogSummary struct {
	Nodes    []HealthSummary
	Services []HealthSummary
	Checks   []HealthSummary
}

type HealthSummary struct {
	Name string `json:",omitempty"`

	Total    int
	Passing  int
	Warning  int
	Critical int

	acl.EnterpriseMeta
}

func (h *HealthSummary) Add(status string) {
	h.Total++
	switch status {
	case api.HealthPassing:
		h.Passing++
	case api.HealthWarning:
		h.Warning++
	case api.HealthCritical:
		h.Critical++
	}
}

type AssignServiceManualVIPsRequest struct {
	Service    string
	ManualVIPs []string

	DCSpecificRequest
}

type AssignServiceManualVIPsResponse struct {
	Found          bool
	UnassignedFrom []PeeredServiceName
}
