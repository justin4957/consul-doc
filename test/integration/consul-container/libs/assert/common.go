// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/assert/common.go
Common module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/assert/common.go> a code:Module ;
    code:name "test/integration/consul-container/libs/assert/common.go" ;
    code:description "Common module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package assert

import (
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	defaultWait    = 500 * time.Millisecond
)
