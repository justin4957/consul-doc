// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/rate/metrics.go
Metrics module for agent layer

## Tags
agent

## Exports
Counters

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/rate/metrics.go> a code:Module ;
    code:name "agent/consul/rate/metrics.go" ;
    code:description "Metrics module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :Counters ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package rate

import "github.com/armon/go-metrics/prometheus"

var Counters = []prometheus.CounterDefinition{
	{
		Name: []string{"rpc", "rate_limit", "exceeded"},
		Help: "Increments whenever an RPC is over a configured rate limit. Note: in permissive mode, the RPC will have still been allowed to proceed.",
	},
	{
		Name: []string{"rpc", "rate_limit", "log_dropped"},
		Help: "Increments whenever a log that is emitted because an RPC exceeded a rate limit gets dropped because the output buffer is full.",
	},
}
