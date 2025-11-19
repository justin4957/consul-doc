// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/bound_refs.go
Bound Refs module for internal layer

## Linked Modules
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
BoundReferenceCollector, NewBoundReferenceCollector

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/bound_refs.go> a code:Module ;
    code:name "internal/resource/bound_refs.go" ;
    code:description "Bound Refs module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :BoundReferenceCollector, :NewBoundReferenceCollector ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"sort"

	"github.com/hashicorp/consul/proto-public/pbresource"
)

type sectionRefKey struct {
	ReferenceKey
	Section string
}

type BoundReferenceCollector struct {
	refs map[sectionRefKey]*pbresource.Reference
}

func NewBoundReferenceCollector() *BoundReferenceCollector {
	return &BoundReferenceCollector{
		refs: make(map[sectionRefKey]*pbresource.Reference),
	}
}

func (c *BoundReferenceCollector) List() []*pbresource.Reference {
	if len(c.refs) == 0 {
		return nil
	}

	out := make([]*pbresource.Reference, 0, len(c.refs))
	for _, ref := range c.refs {
		out = append(out, ref)
	}

	sort.Slice(out, func(i, j int) bool {
		return LessReference(out[i], out[j])
	})

	return out
}

func (c *BoundReferenceCollector) AddRefOrID(ref ReferenceOrID) {
	if c == nil {
		return
	}
	c.AddRef(ReferenceFromReferenceOrID(ref))
}

func (c *BoundReferenceCollector) AddRef(ref *pbresource.Reference) {
	if c == nil {
		return
	}
	srk := sectionRefKey{
		ReferenceKey: NewReferenceKey(ref),
		Section:      ref.Section,
	}

	if _, ok := c.refs[srk]; ok {
		return
	}

	c.refs[srk] = ref
}
