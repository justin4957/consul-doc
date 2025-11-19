// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/operator/raft/operator_raft.go
Operator Raft module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, consensus, replication, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/operator/raft/operator_raft.go> a code:Module ;
    code:name "command/operator/raft/operator_raft.go" ;
    code:description "Operator Raft module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "consensus", "replication", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package raft

import (
	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

const synopsis = "Provides cluster-level tools for Consul operators"
const help = `
Usage: consul operator raft <subcommand> [options]

The Raft operator command is used to interact with Consul's Raft subsystem. The
command can be used to verify Raft peers or in rare cases to recover quorum by
removing invalid peers.
`
