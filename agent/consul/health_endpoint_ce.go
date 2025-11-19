// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/health_endpoint_ce.go
Health Endpoint Ce module for agent layer

## Linked Modules
- [agent/consul/state](../agent/consul/state)
- [agent/structs](../agent/structs)

## Tags
agent, health-checks, monitoring

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/health_endpoint_ce.go> a code:Module ;
    code:name "agent/consul/health_endpoint_ce.go" ;
    code:description "Health Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "health-checks", "monitoring" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"errors"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/structs"
)

// getArgsForSamenessGroupMembers returns the arguments for the sameness group members if SamenessGroup
// field is set in the ServiceSpecificRequest. It returns the index of the sameness group, the arguments
// for the sameness group members and an error if any.
// If SamenessGroup is not set, it returns:
// - the index 0
// - an array containing the original arguments
// - nil error
// If SamenessGroup is set on CE, it returns::
// - the index of 0
// - nil array
// - an error indicating that sameness groups are not supported in consul CE
// If SamenessGroup is set on ENT, it returns:
// - the index of the sameness group
// - an array containing the arguments for the sameness group members
// - nil error
func (h *Health) getArgsForSamenessGroupMembers(args *structs.ServiceSpecificRequest,
	ws memdb.WatchSet, state *state.Store) (uint64, []*structs.ServiceSpecificRequest, error) {
	if args.SamenessGroup != "" {
		return 0, nil, errors.New("sameness groups are not supported in consul CE")
	}
	return 0, []*structs.ServiceSpecificRequest{args}, nil
}
