// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: proto/private/pbautoconf/auto_config_ce.go
Auto Config Ce module for internal layer

## Tags
configuration, internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbautoconf/auto_config_ce.go> a code:Module ;
    code:name "proto/private/pbautoconf/auto_config_ce.go" ;
    code:description "Auto Config Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "configuration", "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbautoconf

func (req *AutoConfigRequest) PartitionOrDefault() string {
	return ""
}
