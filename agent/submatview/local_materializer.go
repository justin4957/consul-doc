// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/submatview/local_materializer.go
Local Materializer module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/consul/state](../agent/consul/state)
- [agent/consul/stream](../agent/consul/stream)
- [lib/retry](../lib/retry)
- [proto/private/pbsubscribe](../proto/private/pbsubscribe)

## Tags
agent

## Exports
ACLResolver, LocalBackend, LocalMaterializer, LocalMaterializerDeps, NewLocalMaterializer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/submatview/local_materializer.go> a code:Module ;
    code:name "agent/submatview/local_materializer.go" ;
    code:description "Local Materializer module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ], [
        code:name "lib/retry" ;
        code:path "../lib/retry"
    ], [
        code:name "proto/private/pbsubscribe" ;
        code:path "../proto/private/pbsubscribe"
    ] ;
    code:exports :ACLResolver, :LocalBackend, :LocalMaterializer, :LocalMaterializerDeps, :NewLocalMaterializer ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package submatview

import (
	"context"
	"errors"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/lib/retry"
	"github.com/hashicorp/consul/proto/private/pbsubscribe"
)

// LocalMaterializer is a materializer for a stream of events
// and manages the local subscription to the event publisher
// until the cache result is discarded when its TTL expires.
type LocalMaterializer struct {
	deps        LocalMaterializerDeps
	retryWaiter *retry.Waiter
	handler     eventHandler

	mat *materializer
}

type LocalMaterializerDeps struct {
	Deps

	Backend     LocalBackend
	ACLResolver ACLResolver
}

var _ Materializer = (*LocalMaterializer)(nil)

type LocalBackend interface {
	Subscribe(req *stream.SubscribeRequest) (*stream.Subscription, error)
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(token string, entMeta *acl.EnterpriseMeta, authzContext *acl.AuthorizerContext) (resolver.Result, error)
}

func NewLocalMaterializer(deps LocalMaterializerDeps) *LocalMaterializer {
	m := LocalMaterializer{
		deps: deps,
		mat:  newMaterializer(deps.Logger, deps.View, deps.Waiter),
	}
	return &m
}

// Query implements Materializer
func (m *LocalMaterializer) Query(ctx context.Context, minIndex uint64) (Result, error) {
	return m.mat.query(ctx, minIndex)
}

// Run receives events from a local subscription backend and sends them to the View.
// It runs until ctx is cancelled, so it is expected to be run in a goroutine.
// Mirrors implementation of RPCMaterializer.
//
// Run implements Materializer
func (m *LocalMaterializer) Run(ctx context.Context) {
	for {
		req := m.deps.Request(m.mat.currentIndex())
		err := m.subscribeOnce(ctx, req)
		if ctx.Err() != nil {
			return
		}
		if m.isTerminalError(err) {
			return
		}

		m.mat.handleError(req, err)

		if err := m.mat.retryWaiter.Wait(ctx); err != nil {
			return
		}
	}
}

// isTerminalError determines whether the given error cannot be recovered from
// and should cause the materializer to halt and be evicted from the view store.
//
// This roughly matches the logic in agent/proxycfg-glue.newUpdateEvent.
func (m *LocalMaterializer) isTerminalError(err error) bool {
	return acl.IsErrNotFound(err)
}

// subscribeOnce opens a new subscription to a local backend and runs
// for its lifetime or until the view is closed.
func (m *LocalMaterializer) subscribeOnce(ctx context.Context, req *pbsubscribe.SubscribeRequest) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	m.handler = initialHandler(req.Index)

	entMeta := req.EnterpriseMeta()
	authz, err := m.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &entMeta, nil)
	if err != nil {
		return err
	}

	subReq, err := state.PBToStreamSubscribeRequest(req, entMeta)
	if err != nil {
		return err
	}

	sub, err := m.deps.Backend.Subscribe(subReq)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	for {
		event, err := sub.Next(ctx)
		switch {
		case errors.Is(err, stream.ErrSubForceClosed):
			m.deps.Logger.Trace("subscription reset by server")
			return err

		case err != nil:
			return err
		}

		if !event.Payload.HasReadPermission(authz) {
			continue
		}

		e := event.Payload.ToSubscriptionEvent(event.Index)

		m.handler, err = m.handler(m, e)
		if err != nil {
			m.mat.reset()
			return err
		}
	}
}

// updateView implements viewState
func (m *LocalMaterializer) updateView(events []*pbsubscribe.Event, index uint64) error {
	return m.mat.updateView(events, index)
}

// reset implements viewState
func (m *LocalMaterializer) reset() {
	m.mat.reset()
}
