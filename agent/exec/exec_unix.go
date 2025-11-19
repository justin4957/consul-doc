// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: agent/exec/exec_unix.go
Exec Unix module for agent layer

## Tags
agent

## Exports
KillCommandSubtree, Script, SetSysProcAttr

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/exec/exec_unix.go> a code:Module ;
    code:name "agent/exec/exec_unix.go" ;
    code:description "Exec Unix module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :KillCommandSubtree, :Script, :SetSysProcAttr ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package exec

import (
	"os"
	"os/exec"
	"syscall"
)

// Script returns a command to execute a script through a shell.
func Script(script string) (*exec.Cmd, error) {
	shell := "/bin/sh"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return exec.Command(shell, "-c", script), nil
}

func SetSysProcAttr(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

// KillCommandSubtree kills the command process and any child processes
func KillCommandSubtree(cmd *exec.Cmd) error {
	return syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
}
