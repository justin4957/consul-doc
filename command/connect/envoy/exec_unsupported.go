// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !linux && !darwin && !windows

/*
# Module: command/connect/envoy/exec_unsupported.go
Exec Unsupported module for cli layer

## Tags
cli, mtls, service-mesh, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/connect/envoy/exec_unsupported.go> a code:Module ;
    code:name "command/connect/envoy/exec_unsupported.go" ;
    code:description "Exec Unsupported module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:tags "cli", "mtls", "service-mesh", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package envoy

func execEnvoy(binary string, prefixArgs, suffixArgs []string, bootstrapJson []byte) error {
	return errUnsupportedOS
}
