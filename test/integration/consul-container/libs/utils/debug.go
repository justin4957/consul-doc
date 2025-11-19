// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/utils/debug.go
Debug module for internal layer

## Tags
internal

## Exports
Dump

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/utils/debug.go> a code:Module ;
    code:name "test/integration/consul-container/libs/utils/debug.go" ;
    code:description "Debug module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Dump ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package utils

import "encoding/json"

// Dump pretty prints the provided arg as json.
func Dump(v any) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "<ERR: " + err.Error() + ">"
	}
	return string(b)
}
