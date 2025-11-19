// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: agent/signal_unix.go
Signal Unix module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/signal_unix.go> a code:Module ;
    code:name "agent/signal_unix.go" ;
    code:description "Signal Unix module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"os"
	"syscall"
)

var forwardSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}
