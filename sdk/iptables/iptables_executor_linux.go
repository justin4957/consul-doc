// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build linux

/*
# Module: sdk/iptables/iptables_executor_linux.go
Iptables Executor Linux module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/iptables/iptables_executor_linux.go> a code:Module ;
    code:name "sdk/iptables/iptables_executor_linux.go" ;
    code:description "Iptables Executor Linux module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package iptables

import (
	"bytes"
	"fmt"
	"os/exec"
)

// iptablesExecutor implements IptablesProvider using exec.Cmd.
type iptablesExecutor struct {
	commands []*exec.Cmd
	cfg      Config
}

func (i *iptablesExecutor) AddRule(name string, args ...string) {
	if i.cfg.NetNS != "" {
		// If network namespace is provided, then we need to execute the command in the given network namespace.
		nsenterArgs := []string{fmt.Sprintf("--net=%s", i.cfg.NetNS), "--", name}
		nsenterArgs = append(nsenterArgs, args...)
		cmd := exec.Command("nsenter", nsenterArgs...)
		i.commands = append(i.commands, cmd)
	} else {
		i.commands = append(i.commands, exec.Command(name, args...))
	}

}

func (i *iptablesExecutor) ApplyRules(command string) error {
	_, err := exec.LookPath(command)
	if err != nil {
		return err
	}

	for _, cmd := range i.commands {
		var cmdOutput bytes.Buffer
		cmd.Stdout = &cmdOutput
		cmd.Stderr = &cmdOutput
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to run command: %s, err: %v, output: %s", cmd.String(), err, cmdOutput.String())
		}
	}

	return nil
}

func (i *iptablesExecutor) Rules() []string {
	var rules []string
	for _, cmd := range i.commands {
		rules = append(rules, cmd.String())
	}

	return rules
}

func (i *iptablesExecutor) ClearAllRules() {
	i.commands = nil
}
