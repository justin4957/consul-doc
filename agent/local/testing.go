// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/local/testing.go
Testing module for agent layer

## Linked Modules
- [agent/token](../agent/token)

## Tags
agent

## Exports
TestState

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/local/testing.go> a code:Module ;
    code:name "agent/local/testing.go" ;
    code:description "Testing module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/token" ;
        code:path "../agent/token"
    ] ;
    code:exports :TestState ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package local

import (
	"os"

	"github.com/hashicorp/consul/agent/token"
	"github.com/hashicorp/go-hclog"
	"github.com/mitchellh/go-testing-interface"
)

// TestState returns a configured *State for testing.
func TestState(_ testing.T) *State {
	logger := hclog.New(&hclog.LoggerOptions{
		Output: os.Stderr,
	})

	result := NewState(Config{}, logger, &token.Store{})
	result.TriggerSyncChanges = func() {}
	return result
}
