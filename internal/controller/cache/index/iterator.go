// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/index/iterator.go
Iterator module for internal layer

## Linked Modules
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/index/iterator.go> a code:Module ;
    code:name "internal/controller/cache/index/iterator.go" ;
    code:description "Iterator module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package index

import "github.com/hashicorp/consul/proto-public/pbresource"

//go:generate mockery --name resourceIterable --with-expecter --exported
type resourceIterable interface {
	Next() ([]byte, []*pbresource.Resource, bool)
}

type resourceIterator struct {
	current []*pbresource.Resource
	iter    resourceIterable
}

func (i *resourceIterator) Next() *pbresource.Resource {
	// Maybe get a new list from the internal iterator
	if len(i.current) == 0 {
		_, i.current, _ = i.iter.Next()
	}

	var rsc *pbresource.Resource
	switch len(i.current) {
	case 0:
		// we are completely out of data so we can return
	case 1:
		rsc = i.current[0]
		i.current = nil
	default:
		rsc = i.current[0]
		i.current = i.current[1:]
	}

	return rsc
}
