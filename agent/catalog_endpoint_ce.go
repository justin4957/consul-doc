// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/catalog_endpoint_ce.go
Catalog Endpoint Ce module for agent layer

## Tags
agent, discovery, registry

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/catalog_endpoint_ce.go> a code:Module ;
    code:name "agent/catalog_endpoint_ce.go" ;
    code:description "Catalog Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "discovery", "registry" .
<!-- End LinkedDoc RDF -->
*/
package agent

import "github.com/armon/go-metrics"

func (s *HTTPHandlers) nodeMetricsLabels() []metrics.Label {
	return []metrics.Label{{Name: "node", Value: s.nodeName()}}
}
