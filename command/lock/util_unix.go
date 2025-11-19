// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: command/lock/util_unix.go
Util Unix module for cli layer

## Tags
cli, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/lock/util_unix.go> a code:Module ;
    code:name "command/lock/util_unix.go" ;
    code:description "Util Unix module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package lock

import (
	"syscall"
)

// signalPid sends a sig signal to the process with process id pid.
func signalPid(pid int, sig syscall.Signal) error {
	return syscall.Kill(pid, sig)
}
