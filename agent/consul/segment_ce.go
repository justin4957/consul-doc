// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/segment_ce.go
Segment Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent

## Exports
SegmentCESummaries

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/segment_ce.go> a code:Module ;
    code:name "agent/consul/segment_ce.go" ;
    code:description "Segment Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :SegmentCESummaries ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"net"

	"github.com/armon/go-metrics/prometheus"

	"github.com/hashicorp/consul/agent/structs"
)

var SegmentCESummaries = []prometheus.SummaryDefinition{
	{
		Name: []string{"leader", "reconcile"},
		Help: "Measures the time spent updating the raft store from the serf member information.",
	},
}

// LANSegmentAddr is used to return the address used for the given LAN segment.
func (s *Server) LANSegmentAddr(name string) string {
	return ""
}

// setupSegmentRPC returns an error if any segments are defined since the CE
// version of Consul doesn't support them.
func (s *Server) setupSegmentRPC() (map[string]net.Listener, error) {
	if len(s.config.Segments) > 0 {
		return nil, structs.ErrSegmentsNotSupported
	}

	return nil, nil
}

// setupSegments returns an error if any segments are defined since the CE
// version of Consul doesn't support them.
func (s *Server) setupSegments(config *Config, rpcListeners map[string]net.Listener) error {
	if len(config.Segments) > 0 {
		return structs.ErrSegmentsNotSupported
	}

	return nil
}

// floodSegments is a NOP in the CE version of Consul.
func (s *Server) floodSegments(config *Config) {
}
