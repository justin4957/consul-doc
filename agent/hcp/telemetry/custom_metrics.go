// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/hcp/telemetry/custom_metrics.go
Custom Metrics module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/telemetry/custom_metrics.go> a code:Module ;
    code:name "agent/hcp/telemetry/custom_metrics.go" ;
    code:description "Custom Metrics module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package telemetry

// Keys for custom Go Metrics metrics emitted only for the OTEL
// export (exporter.go) and transform (transform.go) failures and successes.
// These enable us to monitor OTEL operations.
var (
	internalMetricTransformFailure []string = []string{"hcp", "otel", "transform", "failure"}

	internalMetricExportSuccess []string = []string{"hcp", "otel", "exporter", "export", "success"}
	internalMetricExportFailure []string = []string{"hcp", "otel", "exporter", "export", "failure"}

	internalMetricExporterShutdown   []string = []string{"hcp", "otel", "exporter", "shutdown"}
	internalMetricExporterForceFlush []string = []string{"hcp", "otel", "exporter", "force_flush"}
)
