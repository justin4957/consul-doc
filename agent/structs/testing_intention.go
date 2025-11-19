// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/testing_intention.go
Testing Intention module for agent layer

## Tags
agent, data-model, types

## Exports
TestIntention

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/testing_intention.go> a code:Module ;
    code:name "agent/structs/testing_intention.go" ;
    code:description "Testing Intention module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :TestIntention ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/mitchellh/go-testing-interface"
)

// TestIntention returns a valid, uninserted (no ID set) intention.
func TestIntention(t testing.T) *Intention {
	ixn := &Intention{
		SourceNS:        IntentionDefaultNamespace,
		SourceName:      "api",
		DestinationNS:   IntentionDefaultNamespace,
		DestinationName: "db",
		Action:          IntentionActionAllow,
		SourceType:      IntentionSourceConsul,
		Meta:            map[string]string{},
	}
	ixn.FillPartitionAndNamespace(nil, true)
	return ixn
}
