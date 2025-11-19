// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test-integ/topoutil/http2.go
Http2 module for internal layer

## Tags
internal

## Exports
EnableHTTP2

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test-integ/topoutil/http2.go> a code:Module ;
    code:name "test-integ/topoutil/http2.go" ;
    code:description "Http2 module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :EnableHTTP2 ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package topoutil

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

// EnableHTTP2 returns a new shallow copy of client that has been tweaked to do
// h2c (cleartext http2).
//
// Note that this clears the Client.Transport.Proxy trick because http2 and
// http proxies are incompatible currently in Go.
func EnableHTTP2(client *http.Client) *http.Client {
	// Shallow copy, and swap the transport
	client2 := *client
	client = &client2
	client.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, network, addr)
		},
	}
	return client
}
