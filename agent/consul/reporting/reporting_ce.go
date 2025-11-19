// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/reporting/reporting_ce.go
Reporting Ce module for agent layer

## Tags
agent

## Exports
EntDeps

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/reporting/reporting_ce.go> a code:Module ;
    code:name "agent/consul/reporting/reporting_ce.go" ;
    code:description "Reporting Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EntDeps ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package reporting

import (
	"context"
)

type EntDeps struct{}

func (rm *ReportingManager) initEnterpriseReporting(entDeps EntDeps) error {
	// no op
	return nil
}

func (rm *ReportingManager) StartReportingAgent() error {
	// no op
	return nil
}

func (rm *ReportingManager) StopReportingAgent() error {
	// no op
	return nil
}

func (rm *ReportingManager) RunMetricsWriter(ctx context.Context) {
	// no op
}

func (rm *ReportingManager) RunManualSnapshotWriter(ctx context.Context) {
	// no op
}
