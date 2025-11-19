// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/usagemetrics/usagemetrics_ce.go
Usagemetrics Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/usagemetrics/usagemetrics_ce.go> a code:Module ;
    code:name "agent/consul/usagemetrics/usagemetrics_ce.go" ;
    code:description "Usagemetrics Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package usagemetrics

func newTenancyUsageReporter(u *UsageMetricsReporter) usageReporter {
	return newBaseUsageReporter(u)
}
