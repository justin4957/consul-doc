// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/data_sources_ce.go
Data Sources Ce module for agent layer

## Tags
agent, networking, service-mesh

## Exports
DataSourcesEnterprise

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/data_sources_ce.go> a code:Module ;
    code:name "agent/proxycfg/data_sources_ce.go" ;
    code:description "Data Sources Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :DataSourcesEnterprise ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

type DataSourcesEnterprise struct{}
