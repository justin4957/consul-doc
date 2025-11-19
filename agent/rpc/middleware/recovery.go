// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/rpc/middleware/recovery.go
Recovery module for agent layer

## Tags
agent, communication, networking

## Exports
NewPanicHandler, RecoveryHandlerFunc

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/rpc/middleware/recovery.go> a code:Module ;
    code:name "agent/rpc/middleware/recovery.go" ;
    code:description "Recovery module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :NewPanicHandler, :RecoveryHandlerFunc ;
    code:tags "agent", "communication", "networking" .
<!-- End LinkedDoc RDF -->
*/
package middleware

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
)

// NewPanicHandler returns a RecoveryHandlerFunc type function
// to handle panic in RPC server's handlers.
func NewPanicHandler(logger hclog.Logger) RecoveryHandlerFunc {
	return func(p interface{}) (err error) {
		// Log the panic and the stack trace of the Goroutine that caused the panic.
		stacktrace := hclog.Stacktrace()
		logger.Error("panic serving rpc request",
			"panic", p,
			"stack", stacktrace,
		)

		return fmt.Errorf("rpc: panic serving request")
	}
}

type RecoveryHandlerFunc func(p interface{}) (err error)
