// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows

/*
# Module: command/lock/util_windows.go
Util Windows module for cli layer

## Tags
cli, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/lock/util_windows.go> a code:Module ;
    code:name "command/lock/util_windows.go" ;
    code:description "Util Windows module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package lock

import (
	"os"
	"syscall"
)

// signalPid sends a sig signal to the process with process id pid.
// Since interrupts et al is not implemented on Windows, signalPid
// always sends a SIGKILL signal irrespective of the sig value.
func signalPid(pid int, sig syscall.Signal) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	_ = sig
	return p.Signal(syscall.SIGKILL)
}
