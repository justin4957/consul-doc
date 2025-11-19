// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/exec/exec.go
Exec module for agent layer

## Tags
agent

## Exports
Subprocess

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/exec/exec.go> a code:Module ;
    code:name "agent/exec/exec.go" ;
    code:description "Exec module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :Subprocess ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package exec

import (
	"fmt"
	"os/exec"
)

// Subprocess returns a command to execute a subprocess directly.
func Subprocess(args []string) (*exec.Cmd, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("need an executable to run")
	}
	return exec.Command(args[0], args[1:]...), nil
}
