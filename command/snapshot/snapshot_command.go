// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/snapshot/snapshot_command.go
Snapshot Command module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/snapshot/snapshot_command.go> a code:Module ;
    code:name "command/snapshot/snapshot_command.go" ;
    code:description "Snapshot Command module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package snapshot

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

const synopsis = "Saves, restores and inspects snapshots of Consul server state"
const help = `
Usage: consul snapshot <subcommand> [options] [args]

  This command has subcommands for saving, restoring, and inspecting the state
  of the Consul servers for disaster recovery. These are atomic, point-in-time
  snapshots which include key/value entries, service catalog, prepared queries,
  sessions, and ACLs.

  If ACLs are enabled, a management token must be supplied in order to perform
  snapshot operations.

  Create a snapshot:

      $ consul snapshot save backup.snap

  Restore a snapshot:

      $ consul snapshot restore backup.snap
      
  Decode a snapshot:
  
      $ consul snapshot decode backup.snap

  Inspect a snapshot:

      $ consul snapshot inspect backup.snap

  Run a daemon process that locally saves a snapshot every hour (available only in
  Consul Enterprise) :

      $ consul snapshot agent

  For more examples, ask for subcommand help or view the documentation.
`
