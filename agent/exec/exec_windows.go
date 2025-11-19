// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows

/*
# Module: agent/exec/exec_windows.go
Exec Windows module for agent layer

## Tags
agent

## Exports
KillCommandSubtree, Script, SetSysProcAttr

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/exec/exec_windows.go> a code:Module ;
    code:name "agent/exec/exec_windows.go" ;
    code:description "Exec Windows module for agent layer" ;
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
	"strings"
	"syscall"
)

// Script returns a command to execute a script through a shell.
func Script(script string) (*exec.Cmd, error) {
	shell := "cmd"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	script = "\"" + script + "\""
	cmd := exec.Command(shell, "/C", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CmdLine: strings.Join(cmd.Args, " "),
	}
	return cmd, nil
}

func SetSysProcAttr(cmd *exec.Cmd) {}

func KillCommandSubtree(cmd *exec.Cmd) error {
	return cmd.Process.Kill()
}
