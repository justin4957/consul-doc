// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/operator_backend.go
Operator Backend module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/rpc/operator](../agent/rpc/operator)
- [proto/private/pboperator](../proto/private/pboperator)

## Tags
agent

## Exports
NewOperatorBackend, OperatorBackend

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/operator_backend.go> a code:Module ;
    code:name "agent/consul/operator_backend.go" ;
    code:description "Operator Backend module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/rpc/operator" ;
        code:path "../agent/rpc/operator"
    ], [
        code:name "proto/private/pboperator" ;
        code:path "../proto/private/pboperator"
    ] ;
    code:exports :NewOperatorBackend, :OperatorBackend ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"context"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/rpc/operator"
	"github.com/hashicorp/consul/proto/private/pboperator"
	"github.com/hashicorp/raft"
)

type OperatorBackend struct {
	srv *Server
}

// NewOperatorBackend returns a operator.Backend implementation that is bound to the given server.
func NewOperatorBackend(srv *Server) *OperatorBackend {
	return &OperatorBackend{
		srv: srv,
	}
}

func (op *OperatorBackend) ResolveTokenAndDefaultMeta(token string, entMeta *acl.EnterpriseMeta, authzCtx *acl.AuthorizerContext) (resolver.Result, error) {
	res, err := op.srv.ResolveTokenAndDefaultMeta(token, entMeta, authzCtx)
	if err != nil {
		return resolver.Result{}, err
	}
	if err := op.srv.validateEnterpriseToken(res.ACLIdentity); err != nil {
		return resolver.Result{}, err
	}
	return res, err
}

func (op *OperatorBackend) TransferLeader(_ context.Context, request *pboperator.TransferLeaderRequest) (*pboperator.TransferLeaderResponse, error) {
	reply := new(pboperator.TransferLeaderResponse)
	err := op.srv.attemptLeadershipTransfer(raft.ServerID(request.ID))
	reply.Success = err == nil
	return reply, err
}

var _ operator.Backend = (*OperatorBackend)(nil)
