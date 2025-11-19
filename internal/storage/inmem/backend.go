// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/storage/inmem/backend.go
Backend module for internal layer

## Linked Modules
- [internal/storage](../internal/storage)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
Backend, NewBackend

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/storage/inmem/backend.go> a code:Module ;
    code:name "internal/storage/inmem/backend.go" ;
    code:description "Backend module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/storage" ;
        code:path "../internal/storage"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :Backend, :NewBackend ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package inmem

import (
	"context"
	"strconv"
	"sync/atomic"

	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/internal/storage"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// NewBackend returns a purely in-memory storage backend. It's suitable for
// testing and development mode, but should NOT be used in production as it
// has no support for durable persistence, so all of your data will be lost
// when the process restarts or crashes.
//
// You must call Run before using the backend.
func NewBackend() (*Backend, error) {
	store, err := NewStore()
	if err != nil {
		return nil, err
	}
	return &Backend{store: store}, nil
}

// Backend is a purely in-memory storage backend implementation.
type Backend struct {
	vsn uint64

	store *Store
}

// Run until the given context is canceled. This method blocks, so should be
// called in a goroutine.
func (b *Backend) Run(ctx context.Context) { b.store.Run(ctx) }

// Read implements the storage.Backend interface.
func (b *Backend) Read(_ context.Context, _ storage.ReadConsistency, id *pbresource.ID) (*pbresource.Resource, error) {
	return b.store.Read(id)
}

// WriteCAS implements the storage.Backend interface.
func (b *Backend) WriteCAS(_ context.Context, res *pbresource.Resource) (*pbresource.Resource, error) {
	stored := proto.Clone(res).(*pbresource.Resource)
	stored.Version = strconv.Itoa(int(atomic.AddUint64(&b.vsn, 1)))

	if err := b.store.WriteCAS(stored, res.Version); err != nil {
		return nil, err
	}
	return stored, nil
}

// DeleteCAS implements the storage.Backend interface.
func (b *Backend) DeleteCAS(_ context.Context, id *pbresource.ID, version string) error {
	return b.store.DeleteCAS(id, version)
}

// List implements the storage.Backend interface.
func (b *Backend) List(_ context.Context, _ storage.ReadConsistency, resType storage.UnversionedType, tenancy *pbresource.Tenancy, namePrefix string) ([]*pbresource.Resource, error) {
	return b.store.List(resType, tenancy, namePrefix)
}

// WatchList implements the storage.Backend interface.
func (b *Backend) WatchList(_ context.Context, resType storage.UnversionedType, tenancy *pbresource.Tenancy, namePrefix string) (storage.Watch, error) {
	return b.store.WatchList(resType, tenancy, namePrefix)
}

// ListByOwner implements the storage.Backend interface.
func (b *Backend) ListByOwner(_ context.Context, id *pbresource.ID) ([]*pbresource.Resource, error) {
	return b.store.ListByOwner(id)
}
