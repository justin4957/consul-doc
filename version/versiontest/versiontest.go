// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: version/versiontest/versiontest.go
Versiontest module for internal layer

## Linked Modules
- [version](../version)

## Tags
internal

## Exports
IsEnterprise

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<version/versiontest/versiontest.go> a code:Module ;
    code:name "version/versiontest/versiontest.go" ;
    code:description "Versiontest module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "version" ;
        code:path "../version"
    ] ;
    code:exports :IsEnterprise ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package versiontest

import "github.com/hashicorp/consul/version"

// IsEnterprise returns true if the current build is a Consul Enterprise build.
//
// This should only be called from test code.
func IsEnterprise() bool {
	return version.VersionMetadata == "ent"
}
