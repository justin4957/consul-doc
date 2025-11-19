// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/operations_ce.go
Operations Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/operations_ce.go> a code:Module ;
    code:name "agent/consul/state/operations_ce.go" ;
    code:description "Operations Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
)

func getCompoundWithTxn(tx ReadTxn, table, index string,
	_ *acl.EnterpriseMeta, idxVals ...interface{}) (memdb.ResultIterator, error) {

	return tx.Get(table, index, idxVals...)
}
