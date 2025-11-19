// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !windows

/*
# Module: sdk/freeport/systemlimit.go
Systemlimit module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/freeport/systemlimit.go> a code:Module ;
    code:name "sdk/freeport/systemlimit.go" ;
    code:description "Systemlimit module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package freeport

import "golang.org/x/sys/unix"

func systemLimit() (int, error) {
	var limit unix.Rlimit
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &limit)
	return int(limit.Cur), err
}
