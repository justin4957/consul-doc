// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/testing_ce.go
Testing Ce module for agent layer

## Tags
agent, networking, service-mesh

## Exports
TestDataSourcesEnterprise

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/testing_ce.go> a code:Module ;
    code:name "agent/proxycfg/testing_ce.go" ;
    code:description "Testing Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :TestDataSourcesEnterprise ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

type TestDataSourcesEnterprise struct{}

func (*TestDataSources) buildEnterpriseSources() {}

func (*TestDataSources) fillEnterpriseDataSources(*DataSources) {}

func testConfigSnapshotFixtureEnterprise(*stateConfig) {}
