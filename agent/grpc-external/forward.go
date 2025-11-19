// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/forward.go
Forward module for agent layer

## Tags
agent, api, communication, grpc, networking

## Exports
ForwardMetadataContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/forward.go> a code:Module ;
    code:name "agent/grpc-external/forward.go" ;
    code:description "Forward module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :ForwardMetadataContext ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package external

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func ForwardMetadataContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	return metadata.NewOutgoingContext(ctx, md)
}
