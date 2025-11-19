// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/lease.go
Lease module for internal layer

## Tags
internal

## Exports
Lease

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/lease.go> a code:Module ;
    code:name "internal/controller/lease.go" ;
    code:description "Lease module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Lease ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package controller

// Lease is used to ensure controllers are run as singletons (i.e. one leader-
// elected instance per cluster).
//
// Currently, this is just an abstraction over Raft leadership. In the future,
// we'll build a backend-agnostic leasing system into the Resource Service which
// will allow us to balance controllers between many servers.
type Lease interface {
	// Held returns whether we are the current lease-holders.
	Held() bool

	// Changed returns a channel on which you can receive notifications whenever
	// the lease is acquired or lost.
	Changed() <-chan struct{}
}

type raftLease struct {
	m  *Manager
	ch <-chan struct{}
}

func (l *raftLease) Held() bool               { return l.m.raftLeader.Load() }
func (l *raftLease) Changed() <-chan struct{} { return l.ch }

type eternalLease struct{}

func (eternalLease) Held() bool               { return true }
func (eternalLease) Changed() <-chan struct{} { return nil }
