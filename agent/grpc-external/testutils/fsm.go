// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/testutils/fsm.go
Fsm module for agent layer

## Linked Modules
- [agent/blockingquery](../agent/blockingquery)
- [agent/consul/state](../agent/consul/state)
- [agent/consul/stream](../agent/consul/stream)

## Tags
agent, api, communication, grpc, networking

## Exports
FakeBlockingFSM, FakeFSM, FakeFSMConfig, NewFakeBlockingFSM, Registrar, SetupFSMAndPublisher, TestStateStore

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/testutils/fsm.go> a code:Module ;
    code:name "agent/grpc-external/testutils/fsm.go" ;
    code:description "Fsm module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/blockingquery" ;
        code:path "../agent/blockingquery"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ] ;
    code:exports :FakeBlockingFSM, :FakeFSM, :FakeFSMConfig, :NewFakeBlockingFSM, :Registrar, :SetupFSMAndPublisher, :TestStateStore ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testutils

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/hashicorp/consul/agent/blockingquery"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/stretchr/testify/require"
)

func TestStateStore(t *testing.T, publisher state.EventPublisher) *state.Store {
	t.Helper()

	gc, err := state.NewTombstoneGC(time.Second, time.Millisecond)
	require.NoError(t, err)

	if publisher == nil {
		publisher = stream.NoOpEventPublisher{}
	}

	return state.NewStateStoreWithEventPublisher(gc, publisher)
}

type Registrar func(*FakeFSM, *stream.EventPublisher)

type FakeFSMConfig struct {
	Register  Registrar
	Refresh   []stream.Topic
	publisher *stream.EventPublisher
}

type FakeFSM struct {
	config FakeFSMConfig
	lock   sync.Mutex
	store  *state.Store
}

func newFakeFSM(t *testing.T, config FakeFSMConfig) *FakeFSM {
	t.Helper()

	store := TestStateStore(t, config.publisher)

	fsm := &FakeFSM{store: store, config: config}

	config.Register(fsm, fsm.config.publisher)

	return fsm
}

func (f *FakeFSM) GetStore() *state.Store {
	f.lock.Lock()
	defer f.lock.Unlock()
	return f.store
}

func (f *FakeFSM) ReplaceStore(store *state.Store) {
	f.lock.Lock()
	defer f.lock.Unlock()
	oldStore := f.store
	f.store = store
	oldStore.Abandon()
	for _, topic := range f.config.Refresh {
		f.config.publisher.RefreshTopic(topic)
	}
}

type FakeBlockingFSM struct {
	store *state.Store
}

func NewFakeBlockingFSM(t *testing.T) *FakeBlockingFSM {
	t.Helper()

	store := TestStateStore(t, nil)

	fsm := &FakeBlockingFSM{store: store}

	return fsm
}

func (f *FakeBlockingFSM) GetState() *state.Store {
	return f.store
}

func (f *FakeBlockingFSM) ConsistentRead() error {
	return nil
}

func (f *FakeBlockingFSM) DecrementBlockingQueries() uint64 {
	return 0
}

func (f *FakeBlockingFSM) IncrementBlockingQueries() uint64 {
	return 0
}

func (f *FakeBlockingFSM) GetShutdownChannel() chan struct{} {
	return nil
}

func (f *FakeBlockingFSM) RPCQueryTimeout(queryTimeout time.Duration) time.Duration {
	return queryTimeout
}

func (f *FakeBlockingFSM) SetQueryMeta(blockingquery.ResponseMeta, string) {
}

func SetupFSMAndPublisher(t *testing.T, config FakeFSMConfig) (*FakeFSM, state.EventPublisher) {
	t.Helper()
	config.publisher = stream.NewEventPublisher(10 * time.Second)

	fsm := newFakeFSM(t, config)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go config.publisher.Run(ctx)

	return fsm, config.publisher
}
