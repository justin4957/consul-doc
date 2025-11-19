// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/context.go
Context module for agent layer

## Tags
agent

## Exports
ContextWithRemoteAddr, RemoteAddrFromContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/context.go> a code:Module ;
    code:name "agent/consul/context.go" ;
    code:description "Context module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :ContextWithRemoteAddr, :RemoteAddrFromContext ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"context"
	"net"
)

type contextKeyRemoteAddr struct{}

func ContextWithRemoteAddr(ctx context.Context, addr net.Addr) context.Context {
	return context.WithValue(ctx, contextKeyRemoteAddr{}, addr)
}

func RemoteAddrFromContext(ctx context.Context) (net.Addr, bool) {
	v := ctx.Value(contextKeyRemoteAddr{})
	if v == nil {
		return nil, false
	}
	return v.(net.Addr), true
}
