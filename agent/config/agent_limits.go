// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/config/agent_limits.go
Agent Limits module for agent layer

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/agent_limits.go> a code:Module ;
    code:name "agent/config/agent_limits.go" ;
    code:description "Agent Limits module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

import (
	"fmt"
)

// checkLimitsFromMaxConnsPerClient check that value provided might be OK
// return an error if values are not compatible
func checkLimitsFromMaxConnsPerClient(maxConnsPerClient int) error {
	maxFds, err := getrlimit()
	if err == nil && maxConnsPerClient > 0 {
		// We need the list port + a few at the minimum
		// On Mac OS, 20 FDs are open by Consul without doing anything
		requiredFds := uint64(maxConnsPerClient + 20)
		if maxFds < requiredFds {
			return fmt.Errorf("system allows a max of %d file descriptors, but limits.http_max_conns_per_client: %d needs at least %d", maxFds, maxConnsPerClient, requiredFds)
		}
	}
	return err
}
