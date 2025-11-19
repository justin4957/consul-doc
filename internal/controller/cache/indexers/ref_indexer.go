// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/indexers/ref_indexer.go
Ref Indexer module for internal layer

## Linked Modules
- [internal/controller/cache/index](../internal/controller/cache/index)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
BoundReferences, BoundRefsIndex, RefOrIDFetcher, RefOrIDIndex

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/indexers/ref_indexer.go> a code:Module ;
    code:name "internal/controller/cache/indexers/ref_indexer.go" ;
    code:description "Ref Indexer module for internal layer" ;
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
    code:exports :BoundReferences, :BoundRefsIndex, :RefOrIDFetcher, :RefOrIDIndex ;
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

//go:generate mockery --name RefOrIDFetcher --with-expecter
type RefOrIDFetcher[T proto.Message, V resource.ReferenceOrID] func(*resource.DecodedResource[T]) []V

func RefOrIDIndex[T proto.Message, V resource.ReferenceOrID](name string, fetch RefOrIDFetcher[T, V]) *index.Index {
	return DecodedMultiIndexer[T](name, index.ReferenceOrIDFromArgs, func(r *resource.DecodedResource[T]) (bool, [][]byte, error) {
		refs := fetch(r)
		indexes := make([][]byte, len(refs))
		for idx, ref := range refs {
			indexes[idx] = index.IndexFromRefOrID(ref)
		}

		return len(indexes) > 0, indexes, nil
	})
}

type BoundReferences interface {
	GetBoundReferences() []*pbresource.Reference
	proto.Message
}

func BoundRefsIndex[T BoundReferences](name string) *index.Index {
	return RefOrIDIndex[T](name, func(res *resource.DecodedResource[T]) []*pbresource.Reference {
		return res.Data.GetBoundReferences()
	})
}
