// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/prepared_query_endpoint_ce.go
Prepared Query Endpoint Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/prepared_query_endpoint_ce.go> a code:Module ;
    code:name "agent/consul/prepared_query_endpoint_ce.go" ;
    code:description "Prepared Query Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func parseSameness(svc *structs.ServiceQuery) error {
	if svc.SamenessGroup != "" {
		return fmt.Errorf("sameness-groups are an enterprise-only feature")
	}
	return nil
}

func (sl stateLookup) samenessGroupLookup(_ string, _ acl.EnterpriseMeta) (uint64, *structs.SamenessGroupConfigEntry, error) {
	return 0, nil, nil
}

// GetSamenessGroupFailoverTargets supports Sameness Groups an enterprise only feature. This satisfies the queryServer interface
func (q *queryServerWrapper) GetSamenessGroupFailoverTargets(_ string, _ acl.EnterpriseMeta) ([]structs.QueryFailoverTarget, error) {
	return []structs.QueryFailoverTarget{}, nil
}

func querySameness(_ queryServer,
	_ structs.PreparedQuery,
	_ *structs.PreparedQueryExecuteRequest,
	_ *structs.PreparedQueryExecuteResponse) error {

	return nil
}
