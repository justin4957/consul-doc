// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Package telemetry implements functionality to collect, aggregate, convert and export
// telemetry data in OpenTelemetry Protocol (OTLP) format.
//
// The entrypoint is the OpenTelemetry (OTEL) go-metrics sink which:
// - Receives metric data.
// - Aggregates metric data using the OTEL Go Metrics SDK.
// - Exports metric data using a configurable OTEL exporter.
//
// The package also provides an OTEL exporter implementation to be used within the sink, which:
// - Transforms metric data from the Metrics SDK OTEL representation to OTLP format.
// - Exports OTLP metric data to an external endpoint using a configurable client.

/*
# Module: agent/hcp/telemetry/doc.go
Doc module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/telemetry/doc.go> a code:Module ;
    code:name "agent/hcp/telemetry/doc.go" ;
    code:description "Doc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package telemetry
