// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: sentinel/sentinel_ce.go
Sentinel Ce module for internal layer

## Tags
internal

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sentinel/sentinel_ce.go> a code:Module ;
    code:name "sentinel/sentinel_ce.go" ;
    code:description "Sentinel Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :New ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sentinel

import (
	"github.com/hashicorp/go-hclog"
)

// New returns a new instance of the Sentinel code engine. This is only available
// in Consul Enterprise so this version always returns nil.
func New(logger hclog.Logger) Evaluator {
	return nil
}
