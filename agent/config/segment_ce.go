// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/config/segment_ce.go
Segment Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/segment_ce.go> a code:Module ;
    code:name "agent/config/segment_ce.go" ;
    code:description "Segment Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

import (
	"github.com/hashicorp/consul/agent/structs"
)

func (b *builder) validateSegments(rt RuntimeConfig) error {
	if rt.SegmentName != "" {
		return structs.ErrSegmentsNotSupported
	}
	if len(rt.Segments) > 0 {
		return structs.ErrSegmentsNotSupported
	}
	return nil
}
