// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/indexers/decoded_indexer.go
Decoded Indexer module for internal layer

## Linked Modules
- [internal/controller/cache/index](../internal/controller/cache/index)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
DecodedMultiIndexer, DecodedSingleIndexer, FromArgs, MultiIndexer, SingleIndexer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/indexers/decoded_indexer.go> a code:Module ;
    code:name "internal/controller/cache/indexers/decoded_indexer.go" ;
    code:description "Decoded Indexer module for internal layer" ;
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
    code:exports :DecodedMultiIndexer, :DecodedSingleIndexer, :FromArgs, :MultiIndexer, :SingleIndexer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package indexers

import (
	"github.com/hashicorp/consul/internal/controller/cache/index"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
	"google.golang.org/protobuf/proto"
)

type FromArgs func(args ...any) ([]byte, error)
type SingleIndexer[T proto.Message] func(r *resource.DecodedResource[T]) (bool, []byte, error)
type MultiIndexer[T proto.Message] func(r *resource.DecodedResource[T]) (bool, [][]byte, error)

func DecodedSingleIndexer[T proto.Message](name string, args FromArgs, idx SingleIndexer[T]) *index.Index {
	return index.New(name, &singleIndexer[T]{
		decodedIndexer: idx,
		indexArgs:      args,
	})
}

func DecodedMultiIndexer[T proto.Message](name string, args FromArgs, idx MultiIndexer[T]) *index.Index {
	return index.New(name, &multiIndexer[T]{
		indexArgs:      args,
		decodedIndexer: idx,
	})
}

type singleIndexer[T proto.Message] struct {
	indexArgs      FromArgs
	decodedIndexer SingleIndexer[T]
}

func (i *singleIndexer[T]) FromArgs(args ...any) ([]byte, error) {
	return i.indexArgs(args...)
}

func (i *singleIndexer[T]) FromResource(r *pbresource.Resource) (bool, []byte, error) {
	res, err := resource.Decode[T](r)
	if err != nil {
		return false, nil, err
	}

	return i.decodedIndexer(res)
}

type multiIndexer[T proto.Message] struct {
	decodedIndexer MultiIndexer[T]
	indexArgs      FromArgs
}

func (i *multiIndexer[T]) FromArgs(args ...any) ([]byte, error) {
	return i.indexArgs(args...)
}

func (i *multiIndexer[T]) FromResource(r *pbresource.Resource) (bool, [][]byte, error) {
	res, err := resource.Decode[T](r)
	if err != nil {
		return false, nil, err
	}

	return i.decodedIndexer(res)
}
