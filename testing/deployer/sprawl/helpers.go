// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/helpers.go
Helpers module for internal layer

## Tags
internal

## Exports
IsSquid503, TruncateSquidError

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/helpers.go> a code:Module ;
    code:name "testing/deployer/sprawl/helpers.go" ;
    code:description "Helpers module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :IsSquid503, :TruncateSquidError ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sprawl

// Deprecated: remove
func TruncateSquidError(err error) error {
	return err
}

// Deprecated: remove
func IsSquid503(err error) bool {
	return false
}
