// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: sentinel/evaluator.go
Evaluator module for internal layer

## Tags
internal

## Exports
Evaluator

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sentinel/evaluator.go> a code:Module ;
    code:name "sentinel/evaluator.go" ;
    code:description "Evaluator module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Evaluator ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sentinel

// Evaluator wraps the Sentinel evaluator from the HashiCorp Sentinel policy
// engine.
type Evaluator interface {
	Compile(policy string) error
	Execute(policy string, enforcementLevel string, data map[string]interface{}) bool
	Close()
}
