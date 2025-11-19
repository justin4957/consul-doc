// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: agent/config/limits.go
Limits module for agent layer

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/limits.go> a code:Module ;
    code:name "agent/config/limits.go" ;
    code:description "Limits module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

import "golang.org/x/sys/unix"

// getrlimit return the max file descriptors allocated by system
// return the number of file descriptors max
func getrlimit() (uint64, error) {
	var limit unix.Rlimit
	err := unix.Getrlimit(unix.RLIMIT_NOFILE, &limit)
	// nolint:unconvert // Rlimit.Cur may not be uint64 on all platforms
	return uint64(limit.Cur), err
}
