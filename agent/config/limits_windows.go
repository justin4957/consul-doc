// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows

/*
# Module: agent/config/limits_windows.go
Limits Windows module for agent layer

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/limits_windows.go> a code:Module ;
    code:name "agent/config/limits_windows.go" ;
    code:description "Limits Windows module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

// getrlimit is no-op on Windows, as max fd/process is 2^24 on Wow64
// Return (16 777 216, nil)
func getrlimit() (uint64, error) {
	return 16_777_216, nil
}
