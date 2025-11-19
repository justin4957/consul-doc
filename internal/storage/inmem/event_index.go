// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/storage/inmem/event_index.go
Event Index module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/storage/inmem/event_index.go> a code:Module ;
    code:name "internal/storage/inmem/event_index.go" ;
    code:description "Event Index module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package inmem

import "github.com/hashicorp/go-memdb"

type meta struct {
	Key   string
	Value any
}

func incrementEventIndex(tx *memdb.Txn) (uint64, error) {
	idx, err := currentEventIndex(tx)
	if err != nil {
		return 0, err
	}

	idx++
	if err := tx.Insert(tableNameMetadata, meta{Key: metaKeyEventIndex, Value: idx}); err != nil {
		return 0, nil
	}
	return idx, nil
}

func currentEventIndex(tx *memdb.Txn) (uint64, error) {
	v, err := tx.First(tableNameMetadata, indexNameID, metaKeyEventIndex)
	if err != nil {
		return 0, err
	}
	if v == nil {
		// 0 and 1 index are reserved for special use in the stream package.
		return 2, nil
	}
	return v.(meta).Value.(uint64), nil
}
