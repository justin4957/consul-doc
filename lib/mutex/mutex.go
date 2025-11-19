// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
Package mutex implements the sync.Locker interface using x/sync/semaphore. It
may be used as a replacement for sync.Mutex when one or more goroutines need to
allow their calls to Lock to be cancelled by context cancellation.
*/

/*
# Module: lib/mutex/mutex.go
Mutex module for internal layer

## Tags
internal

## Exports
Mutex, New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/mutex/mutex.go> a code:Module ;
    code:name "lib/mutex/mutex.go" ;
    code:description "Mutex module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Mutex, :New ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package mutex

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Mutex semaphore.Weighted

// New returns a Mutex that is ready for use.
func New() *Mutex {
	return (*Mutex)(semaphore.NewWeighted(1))
}

func (m *Mutex) Lock() {
	_ = (*semaphore.Weighted)(m).Acquire(context.Background(), 1)
}

func (m *Mutex) Unlock() {
	(*semaphore.Weighted)(m).Release(1)
}

// TryLock acquires the mutex, blocking until resources are available or ctx is
// done. On success, returns nil. On failure, returns ctx.Err() and leaves the
// semaphore unchanged.
//
// If ctx is already done, Acquire may still succeed without blocking.
func (m *Mutex) TryLock(ctx context.Context) error {
	return (*semaphore.Weighted)(m).Acquire(ctx, 1)
}
