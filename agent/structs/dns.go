// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/dns.go
Dns module for agent layer

## Tags
agent, data-model, discovery, networking, types

## Exports
RecursorStrategy

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/dns.go> a code:Module ;
    code:name "agent/structs/dns.go" ;
    code:description "Dns module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :RecursorStrategy ;
    code:tags "agent", "data-model", "discovery", "networking", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import "math/rand"

type RecursorStrategy string

const (
	RecursorStrategySequential RecursorStrategy = "sequential"
	RecursorStrategyRandom     RecursorStrategy = "random"
)

func (s RecursorStrategy) Indexes(max int) []int {
	switch s {
	case RecursorStrategyRandom:
		return rand.Perm(max)
	default:
		idxs := make([]int, max)
		for i := range idxs {
			idxs[i] = i
		}
		return idxs

	}
}
