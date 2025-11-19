// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/hcp/deps.go
Deps module for agent layer

## Linked Modules
- [agent/hcp/client](../agent/hcp/client)
- [agent/hcp/config](../agent/hcp/config)
- [agent/hcp/scada](../agent/hcp/scada)
- [agent/hcp/telemetry](../agent/hcp/telemetry)

## Tags
agent

## Exports
Deps, NewDeps

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/deps.go> a code:Module ;
    code:name "agent/hcp/deps.go" ;
    code:description "Deps module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/hcp/client" ;
        code:path "../agent/hcp/client"
    ], [
        code:name "agent/hcp/config" ;
        code:path "../agent/hcp/config"
    ], [
        code:name "agent/hcp/scada" ;
        code:path "../agent/hcp/scada"
    ], [
        code:name "agent/hcp/telemetry" ;
        code:path "../agent/hcp/telemetry"
    ] ;
    code:exports :Deps, :NewDeps ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package hcp

import (
	"context"
	"fmt"

	"github.com/armon/go-metrics"
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel"

	"github.com/hashicorp/consul/agent/hcp/client"
	"github.com/hashicorp/consul/agent/hcp/config"
	"github.com/hashicorp/consul/agent/hcp/scada"
	"github.com/hashicorp/consul/agent/hcp/telemetry"
)

// Deps contains the interfaces that the rest of Consul core depends on for HCP integration.
type Deps struct {
	Config            config.CloudConfig
	Provider          scada.Provider
	Sink              metrics.ShutdownSink
	TelemetryProvider *hcpProviderImpl
	DataDir           string
}

func NewDeps(cfg config.CloudConfig, logger hclog.Logger, dataDir string) (Deps, error) {
	ctx := context.Background()
	ctx = hclog.WithContext(ctx, logger)

	provider, err := scada.New(logger.Named("scada"))
	if err != nil {
		return Deps{}, fmt.Errorf("failed to init scada: %w", err)
	}

	metricsProvider := NewHCPProvider(ctx)
	if err != nil {
		logger.Error("failed to init HCP metrics provider", "error", err)
		return Deps{}, fmt.Errorf("failed to init HCP metrics provider: %w", err)
	}

	metricsClient := client.NewMetricsClient(ctx, metricsProvider)

	sink, err := newSink(ctx, metricsClient, metricsProvider)
	if err != nil {
		// Do not prevent server start if sink init fails, only log error.
		logger.Error("failed to init sink", "error", err)
	}

	return Deps{
		Config:            cfg,
		Provider:          provider,
		Sink:              sink,
		TelemetryProvider: metricsProvider,
		DataDir:           dataDir,
	}, nil
}

// newSink initializes an OTELSink which forwards Consul metrics to HCP.
// This step should not block server initialization, errors are returned, only to be logged.
func newSink(
	ctx context.Context,
	metricsClient telemetry.MetricsClient,
	cfgProvider *hcpProviderImpl,
) (metrics.ShutdownSink, error) {
	logger := hclog.FromContext(ctx)

	// Set the global OTEL error handler. Without this, on any failure to publish metrics in
	// otelExporter.Export, the default OTEL handler logs to stderr without the formatting or group
	// that hclog provides. Here we override that global error handler once so logs are
	// in the standard format and include "hcp" in the group name like:
	// 2024-02-06T22:35:19.072Z [ERROR] agent.hcp: failed to export metrics: failed to export metrics: code 404: 404 page not found
	otel.SetErrorHandler(&otelErrorHandler{logger: logger})

	reader := telemetry.NewOTELReader(metricsClient, cfgProvider)
	sinkOpts := &telemetry.OTELSinkOpts{
		Reader:         reader,
		ConfigProvider: cfgProvider,
	}

	sink, err := telemetry.NewOTELSink(ctx, sinkOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTELSink: %w", err)
	}

	logger.Debug("initialized HCP metrics sink")

	return sink, nil
}

type otelErrorHandler struct {
	logger hclog.Logger
}

func (o *otelErrorHandler) Handle(err error) {
	o.logger.Error(err.Error())
}
