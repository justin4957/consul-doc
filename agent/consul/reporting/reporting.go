// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/reporting/reporting.go
Reporting module for agent layer

## Linked Modules
- [agent/consul/state](../agent/consul/state)
- [agent/structs](../agent/structs)

## Tags
agent

## Exports
NewReportingManager, ReportingManager, ServerDelegate, StateDelegate

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/reporting/reporting.go> a code:Module ;
    code:name "agent/consul/reporting/reporting.go" ;
    code:description "Reporting module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :NewReportingManager, :ReportingManager, :ServerDelegate, :StateDelegate ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package reporting

import (
	"sync"
	"time"

	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-memdb"
	"github.com/hashicorp/go-retryablehttp"
)

type ReportingManager struct {
	logger                      hclog.Logger
	clusterId                   string
	autoReporting               bool
	server                      ServerDelegate
	stateProvider               StateDelegate
	tickerInterval              time.Duration
	manualSnapshotInterval      time.Duration
	snapshotRetention           time.Duration
	initialMetricsRetryInterval time.Duration
	initialMetricsTimeout       time.Duration
	customerID                  string
	EntDeps
	sync.RWMutex
	manualHTTPClient     *retryablehttp.Client
	manualServiceAddress string
}

const (
	SystemMetadataReportingProcessID = "reporting-process-id"
	ReportingInterval                = 1 * time.Hour
)

const (
	// ManualSnapshotInterval controls how often we persist manual census snapshots.
	ManualSnapshotInterval = 24 * time.Hour
	// DefaultSnapshotRetention is the default retention period for manual census snapshots.
	DefaultSnapshotRetention = 9600 * time.Hour // 400 days
)

const (
	defaultInitialMetricsRetryInterval = 500 * time.Millisecond
	defaultInitialMetricsTimeout       = 30 * time.Second
)

//go:generate mockery --name ServerDelegate --inpackage
type ServerDelegate interface {
	GetSystemMetadata(key string) (string, error)
	SetSystemMetadataKey(key, val string) error
	IsLeader() bool
	ApplyCensusRequest(req *structs.CensusRequest) error
}

type StateDelegate interface {
	// Metrics methods
	NodeUsage() (uint64, state.NodeUsage, error)
	ServiceUsage(ws memdb.WatchSet, tenantUsage bool) (uint64, structs.ServiceUsage, error)
	// Census methods
	CensusPut(idx uint64, req *structs.CensusRequest) error
	CensusPrune(idx uint64, cutoff time.Time) (int, error)
	CensusListAll() (uint64, []*state.CensusSnapshot, error)
}

func NewReportingManager(logger hclog.Logger, clusterId string, autoReporting bool, deps EntDeps, server ServerDelegate, stateProvider StateDelegate, snapshotRetention time.Duration) *ReportingManager {
	if snapshotRetention <= 0 {
		snapshotRetention = DefaultSnapshotRetention
	}

	rm := &ReportingManager{
		logger:                      logger.Named("reporting"),
		clusterId:                   clusterId,
		autoReporting:               autoReporting,
		server:                      server,
		stateProvider:               stateProvider,
		tickerInterval:              ReportingInterval,
		manualSnapshotInterval:      ManualSnapshotInterval,
		snapshotRetention:           snapshotRetention,
		initialMetricsRetryInterval: defaultInitialMetricsRetryInterval,
		initialMetricsTimeout:       defaultInitialMetricsTimeout,
	}
	err := rm.initEnterpriseReporting(deps)
	if err != nil {
		rm.logger.Error("Error initializing reporting manager", "error", err)
		return nil
	}
	rm.logger.Debug("Created reporting manager")
	return rm
}

// ConfigureInitialMetricsBootstrap allows callers (primarily tests) to override the retry interval and timeout
// used when warming initial reporting metrics. A timeout <= 0 disables the bootstrap routine entirely.
func (rm *ReportingManager) ConfigureInitialMetricsBootstrap(timeout, interval time.Duration) {
	rm.initialMetricsTimeout = timeout
	rm.initialMetricsRetryInterval = interval
}
