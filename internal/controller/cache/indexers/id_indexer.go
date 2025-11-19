// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/indexers/id_indexer.go
Id Indexer module for internal layer

## Linked Modules
- [internal/controller/cache/index](../internal/controller/cache/index)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
GetSingleRefOrID, IDIndex, OwnerIndex, SingleIDOrRefIndex

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/indexers/id_indexer.go> a code:Module ;
    code:name "internal/controller/cache/indexers/id_indexer.go" ;
    code:description "Id Indexer module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/controller/cache/index" ;
        code:path "../internal/controller/cache/index"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :GetSingleRefOrID, :IDIndex, :OwnerIndex, :SingleIDOrRefIndex ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package indexers

import (
	"github.com/hashicorp/consul/internal/controller/cache/index"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func IDIndex(name string, opts ...index.IndexOption) *index.Index {
	return index.New(name, &idOrRefIndexer{getID: getResourceID}, opts...)
}

func OwnerIndex(name string, opts ...index.IndexOption) *index.Index {
	return index.New(name, &idOrRefIndexer{getID: getOwnerID}, opts...)
}

func SingleIDOrRefIndex(name string, f GetSingleRefOrID, opts ...index.IndexOption) *index.Index {
	return index.New(name, &idOrRefIndexer{getID: f}, opts...)
}

//go:generate mockery --name GetSingleRefOrID --with-expecter
type GetSingleRefOrID func(*pbresource.Resource) resource.ReferenceOrID

func getResourceID(r *pbresource.Resource) resource.ReferenceOrID {
	return r.GetId()
}

func getOwnerID(r *pbresource.Resource) resource.ReferenceOrID {
	return r.GetOwner()
}

type idOrRefIndexer struct {
	getID GetSingleRefOrID
}

// FromArgs constructs a radix tree key from an ID for lookup.
func (i idOrRefIndexer) FromArgs(args ...any) ([]byte, error) {
	return index.MaybePrefixReferenceOrIDFromArgs(args...)
}

// FromObject constructs a radix tree key from a Resource at write-time, or an
// ID at delete-time.
func (i idOrRefIndexer) FromResource(r *pbresource.Resource) (bool, []byte, error) {
	return true, index.IndexFromRefOrID(i.getID(r)), nil
}
