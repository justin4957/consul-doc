// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/watch.go
Watch module for internal layer

## Linked Modules
- [internal/controller/cache/index](../internal/controller/cache/index)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/watch.go> a code:Module ;
    code:name "internal/controller/watch.go" ;
    code:description "Watch module for internal layer" ;
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
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package controller

import (
	"fmt"

	"github.com/hashicorp/consul/internal/controller/cache/index"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

type watch struct {
	watchedType *pbresource.Type
	mapper      DependencyMapper
	indexes     map[string]*index.Index
}

func newWatch(watchedType *pbresource.Type, mapper DependencyMapper) *watch {
	if mapper == nil {
		panic("mapper not provided")
	}

	return &watch{
		watchedType: watchedType,
		indexes:     make(map[string]*index.Index),
		mapper:      mapper,
	}
}

func (w *watch) addIndex(index *index.Index) {
	if _, indexExists := w.indexes[index.Name()]; indexExists {
		panic(fmt.Sprintf("index with name %s is already defined", index.Name()))
	}

	w.indexes[index.Name()] = index
}

func (w *watch) key() string {
	return resource.ToGVK(w.watchedType)
}
