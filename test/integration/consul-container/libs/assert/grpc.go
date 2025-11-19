// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/assert/grpc.go
Grpc module for internal layer

## Linked Modules
- [sdk/testutil/retry](../sdk/testutil/retry)

## Tags
api, communication, grpc, internal, networking

## Exports
GRPCPing

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/assert/grpc.go> a code:Module ;
    code:name "test/integration/consul-container/libs/assert/grpc.go" ;
    code:description "Grpc module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "sdk/testutil/retry" ;
        code:path "../sdk/testutil/retry"
    ] ;
    code:exports :GRPCPing ;
    code:tags "api", "communication", "grpc", "internal", "networking" .
<!-- End LinkedDoc RDF -->
*/
package assert

import (
	"context"
	"testing"
	"time"

	"fortio.org/fortio/fgrpc"
	"github.com/hashicorp/consul/sdk/testutil/retry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCPing sends a fgrpc.PingMessage to a fortio server at addr, analogous to
// the CLI command `fortio grpcping`. It retries for up to 1m, with a 25ms gap.
func GRPCPing(t *testing.T, addr string) {
	t.Helper()
	pingConn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	pingCl := fgrpc.NewPingServerClient(pingConn)
	var msg *fgrpc.PingMessage
	retries := 0
	retry.RunWith(&retry.Timer{Timeout: time.Minute, Wait: 25 * time.Millisecond}, t, func(r *retry.R) {
		r.Logf("making grpc call to %s", addr)
		retries += 1
		msg, err = pingCl.Ping(context.Background(), &fgrpc.PingMessage{
			// use addr as payload so we have something variable to check against
			Payload: addr,
		})
		if err != nil {
			r.Error(err)
		}
	})
	assert.Equal(t, addr, msg.Payload)
}
