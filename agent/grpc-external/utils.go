// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/utils.go
Utils module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)

## Tags
agent, api, communication, grpc, networking

## Exports
ACLResolver, RequireAnyValidACLToken, RequireNotNil, TraceID

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/utils.go> a code:Module ;
    code:name "agent/grpc-external/utils.go" ;
    code:description "Utils module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ] ;
    code:exports :ACLResolver, :RequireAnyValidACLToken, :RequireNotNil, :TraceID ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package external

import (
	"github.com/hashicorp/go-uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
)

// We tag logs with a unique identifier to ease debugging. In the future this
// should probably be a real Open Telemetry trace ID.
func TraceID() string {
	id, err := uuid.GenerateUUID()
	if err != nil {
		return ""
	}
	return id
}

type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

// RequireAnyValidACLToken checks that the caller provided a valid ACL token
// without requiring any specific permissions. This is useful for endpoints
// that are used by all/most consumers of our API, such as those called by the
// consul-server-connection-manager library when establishing a new connection.
//
// Note: no token is required if ACLs are disabled.
func RequireAnyValidACLToken(resolver ACLResolver, token string) error {
	authz, err := resolver.ResolveTokenAndDefaultMeta(token, nil, nil)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	if id := authz.ACLIdentity; id != nil && id.ID() == acl.AnonymousTokenID {
		return status.Error(codes.Unauthenticated, "An ACL token must be provided (via the `x-consul-token` metadata field) to call this endpoint")
	}

	return nil
}

func RequireNotNil(v interface{}, name string) {
	if v == nil {
		panic(name + " is required")
	}
}
