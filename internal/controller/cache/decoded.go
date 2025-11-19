// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/decoded.go
Decoded module for internal layer

## Linked Modules
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
DecodedResourceIterator, GetDecoded, ListDecoded, ListIteratorDecoded, ParentsDecoded, ParentsIteratorDecoded, QueryDecoded

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/decoded.go> a code:Module ;
    code:name "internal/controller/cache/decoded.go" ;
    code:description "Decoded module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :DecodedResourceIterator, :GetDecoded, :ListDecoded, :ListIteratorDecoded, :ParentsDecoded, :ParentsIteratorDecoded, :QueryDecoded ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package cache

import (
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// Get retrieves a single resource from the specified index that matches the provided args.
// If more than one match is found the first is returned.
func GetDecoded[T proto.Message](c ReadOnlyCache, it *pbresource.Type, indexName string, args ...any) (*resource.DecodedResource[T], error) {
	res, err := c.Get(it, indexName, args...)
	if err != nil {
		return nil, err
	}

	if res == nil {
		return nil, nil
	}

	return resource.Decode[T](res)
}

// List retrieves all the resources from the specified index matching the provided args.
func ListDecoded[T proto.Message](c ReadOnlyCache, it *pbresource.Type, indexName string, args ...any) ([]*resource.DecodedResource[T], error) {
	resources, err := c.List(it, indexName, args...)
	if err != nil {
		return nil, err
	}

	return resource.DecodeList[T](resources)
}

// ListIterator retrieves an iterator over all resources from the specified index matching the provided args.
func ListIteratorDecoded[T proto.Message](c ReadOnlyCache, it *pbresource.Type, indexName string, args ...any) (DecodedResourceIterator[T], error) {
	iter, err := c.ListIterator(it, indexName, args...)
	if err != nil {
		return nil, err
	}

	if iter == nil {
		return nil, nil
	}

	return decodedResourceIterator[T]{iter}, nil
}

// Parents retrieves all resources whos index value is a parent (or prefix) of the value calculated
// from the provided args.
func ParentsDecoded[T proto.Message](c ReadOnlyCache, it *pbresource.Type, indexName string, args ...any) ([]*resource.DecodedResource[T], error) {
	resources, err := c.Parents(it, indexName, args...)
	if err != nil {
		return nil, err
	}

	return resource.DecodeList[T](resources)
}

// ParentsIterator retrieves an iterator over all resources whos index value is a parent (or prefix)
// of the value calculated from the provided args.
func ParentsIteratorDecoded[T proto.Message](c ReadOnlyCache, it *pbresource.Type, indexName string, args ...any) (DecodedResourceIterator[T], error) {
	iter, err := c.ParentsIterator(it, indexName, args...)
	if err != nil {
		return nil, err
	}

	if iter == nil {
		return nil, nil
	}

	return decodedResourceIterator[T]{iter}, nil
}

// Query will execute a named query against the cache and return an interator over its results
func QueryDecoded[T proto.Message](c ReadOnlyCache, name string, args ...any) (DecodedResourceIterator[T], error) {
	iter, err := c.Query(name, args...)
	if err != nil {
		return nil, err
	}

	if iter == nil {
		return nil, nil
	}

	return decodedResourceIterator[T]{iter}, nil
}

type DecodedResourceIterator[T proto.Message] interface {
	Next() (*resource.DecodedResource[T], error)
}

type decodedResourceIterator[T proto.Message] struct {
	ResourceIterator
}

func (iter decodedResourceIterator[T]) Next() (*resource.DecodedResource[T], error) {
	res := iter.ResourceIterator.Next()
	if res == nil {
		return nil, nil
	}

	return resource.Decode[T](res)
}
