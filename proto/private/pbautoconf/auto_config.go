// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbautoconf/auto_config.go
Auto Config module for internal layer

## Tags
configuration, internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbautoconf/auto_config.go> a code:Module ;
    code:name "proto/private/pbautoconf/auto_config.go" ;
    code:description "Auto Config module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "configuration", "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbautoconf

import "time"

func (req *AutoConfigRequest) RequestDatacenter() string {
	return req.Datacenter
}

func (req *AutoConfigRequest) IsRead() bool {
	return false
}

func (req *AutoConfigRequest) AllowStaleRead() bool {
	return false
}

func (req *AutoConfigRequest) TokenSecret() string {
	return req.ConsulToken
}

func (req *AutoConfigRequest) SetTokenSecret(token string) {
	req.ConsulToken = token
}

func (req *AutoConfigRequest) HasTimedOut(start time.Time, rpcHoldTimeout, _, _ time.Duration) (bool, error) {
	return time.Since(start) > rpcHoldTimeout, nil
}
