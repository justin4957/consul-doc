// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/debug.go
Debug module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/debug.go> a code:Module ;
    code:name "testing/deployer/sprawl/debug.go" ;
    code:description "Debug module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sprawl

import "encoding/json"

func jd(v any) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
