// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows

/*
# Module: agent/signal_windows.go
Signal Windows module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/signal_windows.go> a code:Module ;
    code:name "agent/signal_windows.go" ;
    code:description "Signal Windows module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"os"
)

var forwardSignals = []os.Signal{os.Interrupt}
