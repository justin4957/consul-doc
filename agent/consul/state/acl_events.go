// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/state/acl_events.go
Acl Events module for agent layer

## Linked Modules
- [agent/consul/stream](../agent/consul/stream)
- [agent/structs](../agent/structs)

## Tags
access-control, agent, authorization, persistence, security, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/acl_events.go> a code:Module ;
    code:name "agent/consul/state/acl_events.go" ;
    code:description "Acl Events module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "access-control", "agent", "authorization", "persistence", "security", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	memdb "github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/agent/structs"
)

// aclChangeUnsubscribeEvent creates and returns stream.UnsubscribeEvents that
// are used to unsubscribe any subscriptions which match the tokens from the events.
//
// These are special events that will never be returned to a subscriber.
func aclChangeUnsubscribeEvent(tx ReadTxn, changes Changes) ([]stream.Event, error) {
	var secretIDs []string

	for _, change := range changes.Changes {
		switch change.Table {
		case tableACLTokens:
			token := changeObject(change).(*structs.ACLToken)
			secretIDs = append(secretIDs, token.SecretID)

		case tableACLRoles:
			role := changeObject(change).(*structs.ACLRole)
			tokens, err := aclTokenListByRole(tx, role.ID, &role.EnterpriseMeta)
			if err != nil {
				return nil, err
			}
			secretIDs = appendSecretIDsFromTokenIterator(secretIDs, tokens)

		case tableACLPolicies:
			policy := changeObject(change).(*structs.ACLPolicy)
			tokens, err := aclTokenListByPolicy(tx, policy.ID, &policy.EnterpriseMeta)
			if err != nil {
				return nil, err
			}
			secretIDs = appendSecretIDsFromTokenIterator(secretIDs, tokens)

			q := Query{Value: policy.ID, EnterpriseMeta: policy.EnterpriseMeta}
			roles, err := tx.Get(tableACLRoles, indexPolicies, q)
			if err != nil {
				return nil, err
			}
			for role := roles.Next(); role != nil; role = roles.Next() {
				role := role.(*structs.ACLRole)

				tokens, err := aclTokenListByRole(tx, role.ID, &policy.EnterpriseMeta)
				if err != nil {
					return nil, err
				}
				secretIDs = appendSecretIDsFromTokenIterator(secretIDs, tokens)
			}
		}
	}
	// There may be duplicate secretIDs here. We rely on this event allowing
	// for duplicate IDs.
	return []stream.Event{stream.NewCloseSubscriptionEvent(secretIDs)}, nil
}

// changeObject returns the object before it was deleted if the change was a delete,
// otherwise returns the object after the change.
func changeObject(change memdb.Change) interface{} {
	if change.Deleted() {
		return change.Before
	}
	return change.After
}

func appendSecretIDsFromTokenIterator(seq []string, tokens memdb.ResultIterator) []string {
	for token := tokens.Next(); token != nil; token = tokens.Next() {
		token := token.(*structs.ACLToken)
		seq = append(seq, token.SecretID)
	}
	return seq
}
